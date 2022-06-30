class TagBlock {
    static Template = document.querySelector("#tmpl-tag-block");

    static Name = TagBlock.Template.content.querySelector(".tmpl-tag-block_name");
    static QuestionsCount = TagBlock.Template.content.querySelector(".tmpl-tag-block_questions-count");

    Name
    QuestionsCount
    
    // GetCopyTagBlock - returns TagBlock as HTMLelement
    static GetCopyElementTagBlock(tag) {
        TagBlock.Name.textContent = tag.Name
        TagBlock.Name.href = `questions/?tag=${tag.Name}`
        TagBlock.QuestionsCount.textContent = tag.QuestionsCount


        return TagBlock.Template.content.cloneNode(true)
    }
}

const B_Tags = document.querySelector(".tags");

//? This is temp solution
function DownloadTags(){
    const tag = new TagBlock()
    tag.Name = "nrblzn"
    tag.QuestionsCount = 100
    
    AppendTagToHTML(tag)
}

function AppendTagToHTML(tag) {
    B_Tags.append(TagBlock.GetCopyElementTagBlock(tag))
}

for (let index = 0; index < 100; index++) {
    DownloadTags()
}