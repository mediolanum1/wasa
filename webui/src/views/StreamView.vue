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
        }
    },


    async mounted() {
        try{ 
     
            const response = await this.$axios.get(`/stream`, 
                 {
                    
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': '   application/json',
                        'Authorization': `Bearer ${this.token}`
                    }, 
                });
                this.photoIds = response.data.reverse()
                
                for (var i = 0; i < response.data.length; i ++){
                    try {
                        const response1 = await this.$axios.get(`/photos/${response.data[i].photoId}`, 
                 {
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${this.token}`,
                    },
                    responseType: 'blob'                
                });
                this.urls.push(URL.createObjectURL(response1.data))
                    }catch(error){

                    }
          
        }
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
                }catch(error){

                }
                this.myFollowings = response1.data
               
                if (this.myFollowings != null){
        
                    if (this.myFollowings.includes(this.userProfile.userId)){          
                     
                        this.followedStatus = true;     
                    }else{
                    this.followedStatus = false;}
                } else{
       
                this.followedStatus = false; }
        
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
            if (this.photoIds[index].likes == null){
                    return false
                }
                if (this.photoIds[index].likes.includes(sessionStorage.getItem('token'))){                
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
            if (this.myFollowings !=null && this.myFollowings.includes(this.userProfile.userId)){
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
                var photoId = this.photoIds[index].photoId
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
                   
                    console.log(' comment deleted successfully', response.data);
                 
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
    <div class="photo-gallery">
     
        <div class="stream" v-for="(photo, index) in urls" :key="index" >
            <div class='singlePost'>
                <div class="photoContainer">
                <img  class='photo' :src="urls[index]">
                </div>
                 <div class="captionPhoto"><h @click="seeAuthor(this.photoIds[index].username)">{{ this.photoIds[index].username }}: </h> <p>  {{this.photoIds[index].photoCaption}}</p></div>
                <div class="photoData">
                    <div class="putLike" @click="toggleLike(this.photoIds[index].photoId, false)"  >
                        <svg  v-if='!checkLike(index)'  class="photoLike"><use href="/feather-sprite-v4.29.0.svg#dislike"></use></svg>
                  
                    </div>

                <div  class="deleteLike" @click="toggleLike(this.photoIds[index].photoId, true)">   <svg  v-if='checkLike(index)'  class="photoLike"><use href="/feather-sprite-v4.29.0.svg#like" style="color:red;"></use></svg></div>
                <div class="likes"><p>likes: {{this.photoIds[index].number_of_likes}}</p></div>
                <div class="date"><p>date: {{this.photoIds[index].photoDate}}</p></div>
                
            
                </div>
  
                <div :id="'commentBox_' + index" class="commentBox"><input placeholder="Add a comment" type="text" class="borderless-input"  :id="index" @input="updateInputValue($event)"></div>
                <div @click='showComments(index)' class="comments"><p>comments: {{this.photoIds[index].number_of_comments}}</p></div>
                
                <!-- Appearing comments block -->
                <div v-if='commentsVisible' class="commentsBox">
                    <!-- Outer part -->
                    <div class="commentsBoxInner">
                        <div class="closeButtonContainer">
                            <button @click="showComments(index)" class="closeButton">X</button>
                        </div>
                        <div class="CommentBox"  v-for="(comment, index1) in photoIds[indexPhoto].comments" :key="index1"> 
                            
                            <div class="singleCommentBox">
                                <button @click="deleteComment(comment.CommentId)"  v-if="checkIfCanDo(comment.UserId)" class="deleteComment">X</button> 
                               <div class="singleComment">{{comment.Comment}}</div> 
                            </div>
                        </div>

                    </div>
                </div>
            </div>
</div>
</div>
</template>
<style>

.stream{
    width: 50%;
    margin: 10px auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background-color: rgb(194, 98, 3);
    align-items: center;

}
.streamData p{
   
    margin: 0;
    margin-left: 5px;
    display: flex;
   
}
.streamData{
    margin-top: 10px;
    display: flex;

}
.photo-gallery {
    width: 100%;
 


  }

.singlePost{
    margin: 0 auto;
    width: 100%;
    background-color: bisque;
    align-items: center;

    justify-content: center;
    display: flex;
    flex-direction: column;
    
}
  .photo-gallery img {
    margin: 5% auto;
    margin-bottom: 20px; 
    
    width: 100%;
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
    
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
.closeButton {
    position: fixed;
    width: 2%;
    height: 5%;
    top: 17%;
    right: 31%;
    
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

.captionPhoto{
    width: 100%;
    margin-left: 50%;
    display: flex;
}
.captionPhoto h{
    font-weight: bold;
    margin-right: 5px;
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

/* Style for profile details */
.profile-details {
    text-align: center;
}

/* Style for user name */
.profile-details h1 {
    font-size: 24px;
    margin-bottom: 10px;
}

/* Style for user details */
.profile-details p {
    font-size: 16px;
    margin-bottom: 5px;
}




  
</style>