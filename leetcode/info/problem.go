package info

import (
	"encoding/json"
	"leetcode-tool/client"
	"leetcode-tool/client/graphql"
)

func GetProblemDetail(titleSlug string, metadataQuery bool) (map[string]json.RawMessage, error) {
	variables := graphql.V{
		"titleSlug": titleSlug,
	}
	query := Query
	if metadataQuery {
		query = QueryMetadata
	}
	reqBody, err := graphql.CreateRequest(query, variables)
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
	if _, exists := result["data"]; !exists {
		return result, nil
	}
	return result, nil
}

const QueryMetadata = `
query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    titleSlug
    judgeType
    exampleTestcases
    codeSnippets {
      langSlug
      code
    }
  }
}
`

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
