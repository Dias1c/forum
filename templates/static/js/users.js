class UserBlock {
    static Template = document.querySelector("#tmpl-user-block");

    static Nickname = UserBlock.Template.content.querySelector(".tmpl-user-block__nickname");
    static BlockAvatarImg = UserBlock.Template.content.querySelector(".tmpl-user-block__block-img");
    static AvatarImg = UserBlock.Template.content.querySelector(".tmpl-user-block__img");
    static CreatedTime = UserBlock.Template.content.querySelector(".tmpl-user-block__created-time");
    static QuestionsCount = UserBlock.Template.content.querySelector(".tmpl-user-block__questions-count");
    static KarmaCount = UserBlock.Template.content.querySelector(".tmpl-user-block__karma-count");

    Name
    AvatarPath
    URL
    CreatedText
    QuestionsCount
    KarmaCount

    // GetCopyUserBlock - returns UserBlock as HTMLelement
    static GetCopyElementUserBlock(user) {
        UserBlock.Nickname.textContent = user.Name
        UserBlock.Nickname.href = user.URL
        UserBlock.BlockAvatarImg.href = user.URL
        UserBlock.AvatarImg.src = user.AvatarPath
        UserBlock.CreatedTime.title = user.CreatedText
        UserBlock.CreatedTime.textContent = user.CreatedText
        UserBlock.QuestionsCount.textContent = user.QuestionsCount
        UserBlock.KarmaCount.textContent = user.KarmaCount

        return UserBlock.Template.content.cloneNode(true)
    }
}

const B_Users = document.querySelector(".users");

//? This is temp solution
function DownloadUsers(){
    const user = new UserBlock()
    user.Name = "Nickname"
    user.AvatarPath = "static/img/avatar.jpeg"
    user.URL = "/"
    user.CreatedText = (new Date()).toISOString().split('T')[0]
    user.QuestionsCount = 1
    user.KarmaCount = 0

    AppendUserToHTML(user)
}

function AppendUserToHTML(user) {
    B_Users.append(UserBlock.GetCopyElementUserBlock(user))
}

for (let index = 0; index < 10; index++) {
    DownloadUsers()
}