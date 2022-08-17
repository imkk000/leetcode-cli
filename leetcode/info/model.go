package info

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

type ResponseBody struct {
	Data struct {
		Question ResponseQuestionBody `json:"question"`
	} `json:"data"`
}

type ResponseQuestionBody struct {
	QuestionId          string                    `json:"questionId"`
	QuestionFrontendId  string                    `json:"questionFrontendId"`
	BoundTopicId        interface{}               `json:"boundTopicId"`
	Title               string                    `json:"title"`
	TitleSlug           string                    `json:"titleSlug"`
	Content             string                    `json:"content"`
	TranslatedTitle     interface{}               `json:"translatedTitle"`
	TranslatedContent   interface{}               `json:"translatedContent"`
	IsPaidOnly          bool                      `json:"isPaidOnly"`
	Difficulty          string                    `json:"difficulty"`
	Likes               int                       `json:"likes"`
	Dislikes            int                       `json:"dislikes"`
	IsLiked             interface{}               `json:"isLiked"`
	RawSimilarQuestions string                    `json:"similarQuestions"`
	SimilarQuestions    []ResponseSimilarQuestion `json:"-"`
	ExampleTestcases    string                    `json:"exampleTestcases"`
	CategoryTitle       string                    `json:"categoryTitle"`
	Contributors        []interface{}             `json:"contributors"`
	TopicTags           []struct {
		Name           string      `json:"name"`
		Slug           string      `json:"slug"`
		TranslatedName interface{} `json:"translatedName"`
		Typename       string      `json:"__typename"`
	} `json:"topicTags"`
	CompanyTagStats interface{} `json:"companyTagStats"`
	CodeSnippets    []struct {
		Lang     string `json:"lang"`
		LangSlug string `json:"langSlug"`
		Code     string `json:"code"`
		Typename string `json:"__typename"`
	} `json:"codeSnippets"`
	RawStats          string               `json:"stats"`
	Stats             ResponseStatsBody    `json:"-"`
	Hints             []string             `json:"hints"`
	Solution          interface{}          `json:"solution"`
	Status            interface{}          `json:"status"`
	SampleTestCase    string               `json:"sampleTestCase"`
	RawMetaData       string               `json:"metaData"`
	MetaData          ResponseMetaDataBody `json:"-"`
	JudgerAvailable   bool                 `json:"judgerAvailable"`
	JudgeType         string               `json:"judgeType"`
	MysqlSchemas      []interface{}        `json:"mysqlSchemas"`
	EnableRunCode     bool                 `json:"enableRunCode"`
	EnableTestMode    bool                 `json:"enableTestMode"`
	EnableDebugger    bool                 `json:"enableDebugger"`
	EnvInfo           string               `json:"envInfo"`
	LibraryUrl        interface{}          `json:"libraryUrl"`
	AdminUrl          interface{}          `json:"adminUrl"`
	ChallengeQuestion interface{}          `json:"challengeQuestion"`
	Typename          string               `json:"__typename"`
}

type ResponseSimilarQuestion struct {
	Title           string      `json:"title"`
	TitleSlug       string      `json:"titleSlug"`
	Difficulty      string      `json:"difficulty"`
	TranslatedTitle interface{} `json:"translatedTitle"`
}

type ResponseStatsBody struct {
	TotalAccepted      string `json:"totalAccepted"`
	TotalSubmission    string `json:"totalSubmission"`
	TotalAcceptedRaw   int    `json:"totalAcceptedRaw"`
	TotalSubmissionRaw int    `json:"totalSubmissionRaw"`
	AcRate             string `json:"acRate"`
}

type ResponseMetaDataBody struct {
	Name        string `json:"name"`
	Classname   string `json:"classname"`
	Constructor struct {
		Params []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"params"`
	} `json:"constructor"`
	Methods []struct {
		Params []struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"params"`
		Name   string `json:"name"`
		Return struct {
			Type string `json:"type"`
		} `json:"return"`
	} `json:"methods"`
	Params []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"params"`
	Return struct {
		Type string `json:"type"`
	} `json:"return"`
	SystemDesign bool `json:"systemdesign"`
}

func (b *Result) UnmarshalJSON(data []byte) error {
	var v ResponseBody
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	q := v.Data.Question
	rawStats := []byte(q.RawStats)
	if err := json.Unmarshal(rawStats, &q.Stats); err != nil {
		return err
	}
	rawSimilarQuestions := []byte(q.RawSimilarQuestions)
	if err := json.Unmarshal(rawSimilarQuestions, &q.SimilarQuestions); err != nil {
		return err
	}
	rawMetadata := []byte(q.RawMetaData)
	if err := json.Unmarshal(rawMetadata, &q.MetaData); err != nil {
		return err
	}

	// set result
	b.Metadata.TopicTags = make([]string, 0)
	for _, topic := range q.TopicTags {
		b.Metadata.TopicTags = append(b.Metadata.TopicTags, topic.Name)
	}
	b.Metadata.SimilarQuestions = make([]ResultMetadataSimilarQuestion, 0)
	for _, sm := range q.SimilarQuestions {
		b.Metadata.SimilarQuestions = append(b.Metadata.SimilarQuestions, ResultMetadataSimilarQuestion{
			Title:      sm.Title,
			TitleSlug:  sm.TitleSlug,
			Difficulty: sm.Difficulty,
		})
	}
	b.Metadata.QuestionID = q.QuestionId
	b.Metadata.QuestionFrontendID = q.QuestionFrontendId
	b.Metadata.Title = q.Title
	b.Metadata.TitleSlug = q.TitleSlug
	b.Metadata.IsPaidOnly = q.IsPaidOnly
	b.Metadata.Difficulty = q.Difficulty
	b.Metadata.Likes = q.Likes
	b.Metadata.Dislikes = q.Dislikes
	b.Metadata.TotalAccepted = q.Stats.TotalAccepted
	b.Metadata.TotalSubmission = q.Stats.TotalSubmission
	b.Metadata.AcRate = q.Stats.AcRate
	b.Metadata.CategoryTitle = q.CategoryTitle

	if q.MetaData.Classname == "" {
		b.Info.ProblemMetadata.Type = "func"
		b.Info.ProblemMetadata.Name = q.MetaData.Name
		params := make([]string, 0)
		for _, param := range q.MetaData.Params {
			params = append(params, fmt.Sprintf("%s:%s", param.Name, param.Type))
		}
		b.Info.ProblemMetadata.Methods = append(b.Info.ProblemMetadata.Methods, ResultInfoProblemMetadataMethod{
			Name:       q.MetaData.Name,
			Params:     params,
			ReturnType: q.MetaData.Return.Type,
		})
	} else {
		b.Info.ProblemMetadata.Type = "class"
		b.Info.ProblemMetadata.Name = q.MetaData.Classname
		b.Info.ProblemMetadata.Methods = make([]ResultInfoProblemMetadataMethod, 0)
		constructorParams := make([]string, 0)
		for _, param := range q.MetaData.Constructor.Params {
			constructorParams = append(constructorParams, fmt.Sprintf("%s:%s", param.Name, param.Type))
		}
		b.Info.ProblemMetadata.Methods = append(b.Info.ProblemMetadata.Methods, ResultInfoProblemMetadataMethod{
			Name:       "constructor",
			Params:     constructorParams,
			ReturnType: q.MetaData.Return.Type,
		})
		for _, method := range q.MetaData.Methods {
			params := make([]string, 0)
			for _, param := range method.Params {
				params = append(params, fmt.Sprintf("%s:%s", param.Name, param.Type))
			}
			b.Info.ProblemMetadata.Methods = append(b.Info.ProblemMetadata.Methods, ResultInfoProblemMetadataMethod{
				Name:       method.Name,
				Params:     params,
				ReturnType: method.Return.Type,
			})
		}
	}
	b.Info.JudgeType = q.JudgeType
	b.Info.ExampleTestcases = strings.Split(q.ExampleTestcases, "\n")
	b.Info.RawExampleTestcases = q.ExampleTestcases
	b.Info.Hints = q.Hints

	return nil
}

type Result struct {
	Metadata ResultMetadata `json:"metadata"`
	Info     ResultInfo     `json:"info"`
}

type ResultInfo struct {
	ProblemMetadata     ResultInfoProblemMetadata `json:"problem_metadata"`
	JudgeType           string                    `json:"judge_type"`
	RawExampleTestcases string                    `json:"raw_example_testcases"`
	ExampleTestcases    []string                  `json:"example_testcases"`
	Hints               []string                  `json:"hints"`
	LanguageSnippets    map[string]string         `json:"language_snippets"`
}

type ResultInfoProblemMetadata struct {
	Type    string                            `json:"type"`
	Name    string                            `json:"name"`
	Methods []ResultInfoProblemMetadataMethod `json:"methods"`
}

type ResultInfoProblemMetadataMethod struct {
	Name       string   `json:"name"`
	Params     []string `json:"params"`
	ReturnType string   `json:"return_type"`
}

type ResultMetadataSimilarQuestion struct {
	Title      string `json:"title"`
	TitleSlug  string `json:"title_slug"`
	Difficulty string `json:"difficulty"`
}

type ResultMetadata struct {
	QuestionID         string                          `json:"question_id"`
	QuestionFrontendID string                          `json:"question_frontend_id"`
	Title              string                          `json:"title"`
	TitleSlug          string                          `json:"title_slug"`
	IsPaidOnly         bool                            `json:"is_paid_only"`
	Difficulty         string                          `json:"difficulty"`
	TotalAccepted      string                          `json:"total_accepted"`
	TotalSubmission    string                          `json:"total_submission"`
	AcRate             string                          `json:"ac_rate"`
	Likes              int                             `json:"likes"`
	Dislikes           int                             `json:"dislikes"`
	CategoryTitle      string                          `json:"category_title"`
	TopicTags          []string                        `json:"topic_tags"`
	SimilarQuestions   []ResultMetadataSimilarQuestion `json:"similar_questions"`
}
