package info

import (
	"encoding/json"
	"leetcode-tool/client"
	"leetcode-tool/client/graphql"
)

func GetProblemDetail(titleSlug string) (map[string]json.RawMessage, error) {
	variables := graphql.V{
		"titleSlug": titleSlug,
	}
	reqBody, err := graphql.CreateRequest(Query, variables)
	if err != nil {
		return nil, err
	}
	resp, err := client.Post("graphql", reqBody)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var result map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

const Query = `
query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    boundTopicId
    title
    titleSlug
    content
    translatedTitle
    translatedContent
    isPaidOnly
    difficulty
    likes
    dislikes
    isLiked
    similarQuestions
    exampleTestcases
    categoryTitle
    contributors {
      username
      profileUrl
      avatarUrl
      __typename
    }
    topicTags {
      name
      slug
      translatedName
      __typename
    }
    companyTagStats
    codeSnippets {
      lang
      langSlug
      code
      __typename
    }
    stats
    hints
    solution {
      id
      canSeeDetail
      paidOnly
      hasVideoSolution
      paidOnlyVideo
      __typename
    }
    status
    sampleTestCase
    metaData
    judgerAvailable
    judgeType
    mysqlSchemas
    enableRunCode
    enableTestMode
    enableDebugger
    envInfo
    libraryUrl
    adminUrl
    challengeQuestion {
      id
      date
      incompleteChallengeCount
      streakCount
      type
      __typename
    }
    __typename
  }
}
`
