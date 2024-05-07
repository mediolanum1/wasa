<script>

export default {
    data(){ 
        return {
            username:  sessionStorage.getItem('username'),
            token: sessionStorage.getItem('token'),
            userProfile: 0,
            urls: [],
            commentsVisible: false,
            indexPhoto: 0,
            show: false,
            showFollowers: false,
            followedStatus: null,
            bannedStatus: null,
            showFollowings: false,
            followings: [],
            followers: [],
        }
    },

    watch:  {
        '$route.params.username'(newParam, oldParam) {
            if (newParam !== oldParam) {
                location.reload();
            }
        },
    },

    async mounted() {
        try{ 
     
            const response = await this.$axios.get(`/users/`, 
                 {
                    params: {
                        'username': this.$route.params.username
                    },
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': '   application/json',
                        'Authorization': `Bearer ${this.token}`
                    }, 
                });
                console.log(response.data)
                this.userProfile = response.data  
                this.checkFollow()
                this.bannedStatus = this.checkBan()
                this.getPhotoBody()
                this.getFollowers();
                this.getFollowings();
                
            } catch(error){
                console.log(error)
            }
        },

        
    methods: {

        async getPhotoBody(){ 
            for (var i = 0; i < this.userProfile.photos.length; i ++){
                    try {
                        const response = await this.$axios.get(`/photos/${this.userProfile.photos[i].photoId}`, 
                    { 
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${this.token}`,
                    },
                    responseType: 'blob'                
                });
                this.urls.push(URL.createObjectURL(response.data))
      
                    } catch(error){ 
                        console.log(error)
                    }
        }
      },

        checkIfCanDo(authorId){
              
                if (this.token == authorId){
                    return true
                }
                return false
        },

        async checkFollow(){
      
            try {
                var response1 = await  this.$axios.get(`/followings/${this.token}`, 
                 {           
                    headers: {            
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${this.token}`,
                    },   
                                   
                });
                this.myFollowings = response1.data
                console.log("this is my followigns",response1.data)
                if (this.myFollowings != null){
        
                    if (this.myFollowings.includes(this.userProfile.username)){          
                     
                        this.followedStatus = true;     
                    }else{
                    this.followedStatus = false;}
                } else{
       
                this.followedStatus = false; }
                }catch(error){

                }

                
        
        },

        async checkBan(){   
          
            try {
                var response1 = await  this.$axios.get(`/banned/${this.token}`, 
                 {           
                    headers: {            
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${this.token}`,
                    },                  
                });
                }catch(error){

                }
            
                this.myBanned = response1.data
             
                if (this.myBanned != null){   
                    if (this.myBanned.includes(this.userProfile.userId)){          
                        this.bannedStatus = false;
                        return true       
                    }
                    return false
                }  
                return false
        },

        checkLike(index){
            if (this.userProfile.photos[index].likes == null){
                    return false
                }
                if (this.userProfile.photos[index].likes.includes(sessionStorage.getItem('token'))){                
                    return true
                } 
                return false
        },
        showComments(index){
                this.indexPhoto = index
                this.commentsVisible = !this.commentsVisible
        },

        async toggleLike(photoId, toggle){
            if (!toggle){
            try {
                const response = await this.$axios.put(`/photos/${photoId}/likes/self`, {},{
                    headers: {
                          Authorization: `Bearer ${this.token}`,
                        'Content-Type': 'application/json',
                    },
                });
                   
                    console.log('Photo liked successfully', response.data);
                
                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
            } else{ 
                try {
                const response = await this.$axios.delete(`/photos/${photoId}/likes/self`,{
                    headers: {
                        Authorization: `Bearer ${this.token}`,
                        'Content-Type': 'application/json',
                    },
                });
                    console.log('Photo unliked successfully', response.data);            
                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
            }
        },
        
        async toggleFollow(){
            console.log("this areeeee my followings",this.myFollowings)
            if (this.myFollowings !=null && this.myFollowings.includes(this.userProfile.username)){
                try {
                   const response = await this.$axios.delete(`/followings/${this.userProfile.userId}`,{
                       headers: {
                           Authorization: `Bearer ${this.token}`,
                           'Content-Type': 'application/json',
                       },
                   });
                       console.log('Unfollowed successfully', response.data);
                       location.reload()
                   }
                   catch (error) {
                       const statusCode = error.response;
                       switch (statusCode) {
                       case 401:
                           console.error('Access Unauthorized');
                           this.photoUploadResult = "You have to log in to post a photo";
                           this.uploadSuccess = true;
                           break;
                       default:
                           console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                       }
                   }
            }else{
                try {
                const response = await this.$axios.put(`/followings/${this.userProfile.userId}`, {},{
                    headers: {
                          Authorization: `Bearer ${this.token}`,
                        'Content-Type': 'application/json',
                    },
                });
                 
                    console.log('followed user successfully', response.data);

                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
                
            }

        },

        async toggleBan(){
            
            if (this.myBanned!=null && this.myBanned.includes(this.userProfile.userId)){
                try {
                   const response = await this.$axios.delete(`/banned/${this.userProfile.userId}`,{
                       headers: {
                           Authorization: `Bearer ${this.token}`,
                           'Content-Type': 'application/json',
                       },
                   });
                       console.log('Unbanned successfully', response.data);
                       location.reload()
                   }
                   catch (error) {
                       const statusCode = error.response;
                       switch (statusCode) {
                       case 401:
                           console.error('Access Unauthorized');
                           this.photoUploadResult = "You have to log in to post a photo";
                           this.uploadSuccess = true;
                           break;
                       default:
                           console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                       }
                   }
            }else{
                try {

                this.toggleFollow()
                const response = await this.$axios.put(`/banned/${this.userProfile.userId}`, {},{
                    headers: {
                          Authorization: `Bearer ${this.token}`,
                        'Content-Type': 'application/json',
                    },
                });
                 
                    console.log('Banned user successfully', response.data);

                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
                
            }
        },

        updateInputValue(event){
              
                var index =event.target.id
              
                var element =  document.getElementById('commentBox_' + index);  
                if (this.comment != ""){
                    if (!element.querySelector('button')){
                        var button = document.createElement('button');
                        button.textContent = 'Post';
                        button.addEventListener('click',  () => this.postComment(index, event.target.value))
                        element.appendChild(button);
                        
                    }
            }
                else{
                   var toRemove = element.getElementsByTagName('button')[0];
                    element.removeChild(toRemove);
                }
        },

        async postComment(index, comment){
                var photoId = this.userProfile.photos[index].photoId
                try {
                const response = await this.$axios.post(`/photos/${photoId}/comments`,{
                        'comment': comment
                },{
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
                        'Content-Type': 'application/json',
                    },
                   
                });
          
                    console.log('Comment uploaded successfully', response.data);
                 
                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a comment";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
        },

        async deleteComment(comment){
       
                try {
                const response = await this.$axios.delete(`/comments/${comment}`,{
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
                        'Content-Type': 'application/json',
                    },
                });
                   
                    console.log(' Photo deleted successfully', response.data);
                 
                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                        break;
                    }
                }
        },

   
        async updateUsername(){
            try {
                const response =  await this.$axios.put(`/username`, {'username': this.newUsername},{
                    headers: {
                          Authorization: `Bearer ${this.token}`,
                        'Content-Type': 'application/json',
                    },
                });
                console.log('Username updated successfully', response.data);
                sessionStorage.setItem('username', this.newUsername)
                    this.username  = this.newUsername
                    this.$router.push('/users/' + sessionStorage.getItem('username'));   
                    location.reload()
                
                }
                
                catch (error) {
                    
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    case 400: 
                        console.error("username already exists")
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                        break;
                        
                    }
                  
                }
        },

        async getFollowers(){
            try {
                const response = await this.$axios.get(`/followers/${this.token}`,{
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
                        'Content-Type': 'application/json',
                    },
                });
                   
                    console.log(' returned followers successfully', response.data);
                    this.followers = response.data
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
        },

        async getFollowings(){ 
            try {
                const response = await this.$axios.get(`/followings/${this.token}`,{
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
                        'Content-Type': 'application/json',
                    },
                });
                   
                    console.log(' returned followings successfully', response.data);
                    this.followings = response.data
                  
                 
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
           
        },
        seeAuthor(username){
                this.$router.push('/users/' + username);  
                
            },

        async deletePhoto(photoId){
            try {
                const response = await this.$axios.delete(`/photos/${photoId}`,{
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
                        'Content-Type': 'application/json',
                    },
                });
                   
                    console.log(' Photo deleted successfully', response.data);
                 
                    location.reload()
                }
                catch (error) {
                    const statusCode = error.response;
                    switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized');
                        this.photoUploadResult = "You have to log in to post a photo";
                        this.uploadSuccess = true;
                        break;
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error);
                    }
                }
        }






            }
}
    

</script>


<template>
    <div class="profile">
        <div class="profile-container">
            <div class="profile-picture">
                <!-- Placeholder image for profile picture -->
                <div v-if="this.userProfile.number_of_photos > 0">
                    <img :src="urls[0]"  alt="Profile pic">
                </div>
                <div v-else>
                    <svg class="icon"><use href="/feather-sprite-v4.29.0.svg#meh"></use></svg>
                 
                </div>
            </div>
            <div class="profile-details">
              
                <p style="font-weight:bold; font-size:x-large" @click="() => { if (this.show = checkIfCanDo(this.userProfile.userId) ? true: false  ); }">{{ this.$route.params.username }}</p>
                <div v-if="this.show" class="changeUsername"><p>change username:</p>
                    <input placeholder="Your new username" type="text" class="borderless-input"  v-model="newUsername" >
                    <button @click="updateUsername" type="submit">Submit</button>
                </div>
                <p>Photos: {{ this.userProfile.number_of_photos}}</p>

                <p @click="() => this.showFollowers=true"> Followers: {{ this.userProfile.number_of_followers }}</p>
                <div v-if="showFollowers" class="changeUsername"><p>Followers: </p>
                    <div class="" >
                        <button @click="() =>  this.showFollowers= false " class="">X</button>
                    </div>
                    <div  v-for="follower in this.followers" :key="follower.userId">
                        
                    <p @click="seeAuthor(follower)" >{{follower}}</p>
                </div>
                </div>
                <p @click="() => this.showFollowings=true">Followings: {{ this.userProfile.number_of_followings}}</p>
                <div v-if="this.showFollowings" class="changeUsername"><p>Followings: </p>
                    <div class="" >
                        <button @click="() =>  this.showFollowings= false " class="">X</button>
                    </div>
                    <div  v-for="follower in this.followings" :key="follower.userId">
                        
                    <p @click="seeAuthor(follower)" >{{follower}}</p>
                </div>
                </div>
            </div>
            <div class='follow-unfollow-ban'>
                <button  v-if="this.userProfile.userId != this.token"  @click='toggleFollow' >  {{ followedStatus ? 'Unfollow' : 'Follow' }}</button> 
                <button @click='toggleBan' v-if="this.userProfile.userId != this.token"> {{ bannedStatus ? 'Ban' : 'Unban' }}</button>
            
            </div>
            </div>
           
                <div class="row">
                    <div class="photos" v-for="( _ , index) in urls" :key="index" >
                        <button class="deletePhoto"  v-if="checkIfCanDo(this.userProfile.photos[index].userId)" @click="deletePhoto(this.userProfile.photos[index].photoId)" >X</button>
                            <div class="photoContainer">
                            <img  class='photo' :src="urls[index]">
                            </div>
                             <div class="caption"><p style="font-weight:bold">caption:</p><p> {{this.userProfile.photos[index].photoCaption}}</p></div>
                            <div class="photoData">
                                <div class="putLike" @click="toggleLike(this.userProfile.photos[index].photoId,checkLike(index) )"  >
                                    <svg  v-if='!checkLike(index)'  class="photoLike"><use href="/feather-sprite-v4.29.0.svg#dislike"></use></svg>
                                </div>
                            <div  class="deleteLike" @click="toggleLike(this.userProfile.photos[index].photoId, checkLike(index))"><svg  v-if='checkLike(index)'  class="photoLike"><use href="/feather-sprite-v4.29.0.svg#like"></use></svg></div>
                            <div class="likes"><p>likes: {{this.userProfile.photos[index].number_of_likes}}</p></div>
                            <div class="date"><p>date: {{this.userProfile.photos[index].photoDate}}</p></div>
                            
                        
                            </div>
              
                            <div :id="'commentBox_' + index" class="commentBox"><input placeholder="Add a comment" type="text" class="borderless-input"  :id="index" @input="updateInputValue($event)"></div>
                            <div @click='showComments(index)' class="comments"><p>comments: {{this.userProfile.photos[index].number_of_comments}}</p></div>
                            
                            <!-- Appearing comments block -->
                            <div v-if='commentsVisible' class="commentsBox">
                                <!-- Outer part -->
                                <div class="commentsBoxInner">
                                    <div class="closeButtonContainer">
                                        <button @click="showComments(index)" class="closeButton">X</button>
                                    </div>
                                    <div class="CommentBox"  v-for="(comment, index1) in this.userProfile.photos[indexPhoto].comments" :key="index1"> 
                                        <div class="singleCommentBox">
                                            <button @click="deleteComment(comment.CommentId)"  v-if="checkIfCanDo(comment.UserId)" class="deleteComment">X</button> 
                                           <p class="singleComment">{{comment.Comment}}</p> 
                                        </div>
                                    </div>
    
                                </div>
                            </div>
            
        </div>
    </div>
    </div>
</template>




<style>

.changeUsername{
    width: 300px;
    height: 300px;
    position: fixed;
    top: 100px;
    left: 40%;
    background-color: antiquewhite;
    font-size: larger;
    font-weight: 600;
}

.deleteComment{
    opacity: 0;
    width: 30px;
    height: 30px;
    position: relative;
    float: right;
    top: 7.5%;
    background-color: #eb4e4e;
    font-weight: 900;
    border-radius: 5px;
    
}

.deleteComment:hover{
    opacity: 1;
   

    
}

.deletePhoto{
    opacity: 0;
    width: 40px;
    height: 40px;
    position: relative;
    left: 51%;
    top: 7.5%;
    background-color: #eb4e4e;
    font-weight: 900;
    


}
.deletePhoto:hover{
    opacity: 1;
   

    
}
.CommentBox{
    width: 100%;
    margin: 10px;
    
    flex-direction: column;


}
.singleCommentBox{
    height: auto;
    min-height: 40px;
    background-color: #a2a2a2; 
    border: 1px solid black;

}
.singleComment{
    padding-top:10px ;
    margin-left: 10px;

}

::-webkit-scrollbar-track {
    background: rgba(0,4,53,0.1); 
}

.closeButtonContainer {
    position: relative;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    
  }
  
.closeButton {
    
    width: 30px;
    height: 30px;
    
    margin: 0;
    background-color: #0e0e0e; 
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 14px;
  }
  
.closeButton:hover {
    background-color: #cc0000; 
  }
.commentsBox{
   
    position: fixed;
    z-index: 1;
    left: 30%;
    top: 15%;
    width: 40%;
    height: 60%;
    overflow: auto;
    background-color: rgb(0, 0, 255);
    background-color: rgb(160, 160, 168);
    border-radius: 10px;
  
}
::-webkit-scrollbar {
    width:0; 
    height: 0; 
}


.commentsBoxInner{
  
    flex-wrap: wrap;
    overflow-y: scroll;
    display: flex;
    position: fixed;
    border-radius: 5px;
    z-index: 1;
    left: 31%;
    top: 17%;
    width: 38%;
    height: 56%;
    overflow: auto;
    background-color: rgb(0, 0, 255);
    background-color: rgb(237, 237, 237);


}
.commentsBoxInner h4{
    
    margin: 0 auto;
}
.profile{
    width: 100%;
}
.follow-unfollow-ban{
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    width: 20%;
    margin-top: 1%;
    padding-bottom: 3%;
  

}
.follow-unfollow-ban button{
border-radius: 5px;

background-color: rgb(209, 116, 239);
border: solid 1px  rgba(0, 0, 0, 0.1) ;
width: 30%;
margin-left: 20px;
margin-right: 20px;
}

.commentBox button{
    width: auto;
    height: 40%;
    background-color: rgb(255, 255, 255);
    color: rgb(212, 91, 246);
    border: 1px solid rgba(0, 0, 0, 0.1);
    font-weight: bold;
}
.commentBox button:hover{
    width: auto;
    height: 40%;
    background-color: rgb(111, 10, 169);
    color: rgb(212, 91, 246);
    border: 1px solid rgba(0, 0, 0, 0.1);
    font-weight: bold;
}

.caption h{
    font-weight: bold;
}

.photoLike{
    width: 3vw;
    height: 5vh;
    padding-right: 10px;
}
.borderless-input {
    border: none; 

    outline: none; 
    background-color: transparent; 
    width: 70%; 
    padding-bottom: 10px;
  }
.row{
    margin-left: 10%;

    display: grid;
    grid-template-columns: repeat(3, 1fr); 
    
}


.photoContainer{
    width: 200px; 
    height: 200px; 
    overflow: hidden; 
}
.photoContainer img{
    width: 100%;
    height: 100%;
    object-fit: cover;  
}
.date{ 
    margin: auto;
    margin-left: 40px; 
}
.photoData{
    display: flex;
    justify-content: space-between;
}

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}


.profile-container {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin-top: 50px;
}


.profile-picture img {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    object-fit: cover;
    margin-bottom: 20px;
}


.profile-details {
    text-align: center;
}

.profile-details h1 {
    font-size: 24px;
    margin-bottom: 10px;
}

.profile-details p {
    font-size: 16px;
    margin-bottom: 5px;
}


</style>