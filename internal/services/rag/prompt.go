package rag

import "fmt"

const RAGSystemPrompt = `
You are an AI assistant that helps users manage their job applications.

Your responsibilities:
- Answer questions ONLY using the provided context.
- If the answer is not present in the context, say that you don't know.
- Never make up companies, interview dates, meeting links, or application statuses.
- Be concise and accurate.
- If the user asks for a meeting link, test link, interview date, next step, comments, or application status, use only the information present in the context.
- Format answers naturally and clearly.
`

func BuildRAGUserPrompt(question, context string) string {
	return fmt.Sprintf(`
Context:
%s

Question:
%s

Answer the question using only the information from the context.
If the context does not contain enough information, say that you don't know.
`, context, question)
}
