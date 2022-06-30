// Widgets like Search

// HEADER SEARCH
var cb_QuestionsFilter = document.getElementById("cb_QuestionsFilter");
var b_QuestionsFilter = document.getElementById("b_QuestionsFilter");


// cb_HeaderSearch_change - Shows or hide element
function cb_QuestionsFilter_change(e) {
    if (e.target.checked) {
        b_QuestionsFilter.classList.add("d-flex");
    } else {
        b_QuestionsFilter.classList.remove("d-flex");
    }
}
cb_QuestionsFilter.addEventListener("change", cb_QuestionsFilter_change);

// Download Questions
class QuestionBlock {
    static Template = document.querySelector("#tmpl-question-block");

    // static Name = TagBlock.Template.content.querySelector(".tmpl-tag-block_name");

    Title
    Text
    Tags
    VotesCount
    AnswersCount
    IsResolved
    CreatedTime
    URL
    AuthorName
    AuthorURL

    // GetCopyQuestionBlock - returns QuestionBlock as HTMLelement
    static GetCopyElementQuestionBlock(quesiton) {
        // Title
        // Text
        // Tags
        // KarmaCount
        // AnswersCount
        // IsResolved
        // CreatedTime
        // URL
        // AuthorName
        // AuthorURL

        return QuestionBlock.Template.content.cloneNode(true)
    }
}

const B_Questions = document.querySelector(".questions");

//? This is temp solution
function DownloadQuestions(){
    const question = new QuestionBlock()
    question.Title = "How can I pass the `Graphql` mutation query in cypress req and get the data back in response?";
    question.Text = "How can I pass the Graphql mutation query in cypress req and get the data back in response ? I am getting an error like Validation error of type FieldUndefined: Field 'mutation' in type Query undefined @muation";
    question.Tags = ["graphql", "cypress"]
    
    AppendQuestionToHTML(question)
}

function AppendQuestionToHTML(question) {
    B_Questions.append(QuestionBlock.GetCopyElementQuestionBlock(question))
}

for (let index = 0; index < 1; index++) {
    DownloadQuestions()
}