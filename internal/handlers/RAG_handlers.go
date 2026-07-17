package handlers


func RAG(repo *repository.OpportunityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


				answer, err := ragService.Answer(question)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		
}