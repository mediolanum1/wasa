<script>
    export default{
        data(){
            return {
                photo: null,
                caption: '',
                buttonText: 'Upload',
                disabled: false,
            }

        },
        methods: {  
            updatePhotoValue(event) {
                this.photo = event.target.files[0]; // Update input value as the user types
            },
            updateCaptionValue(event) {
                this.caption = event.target.value; // Update input value as the user types
            },
           async upload(){
            const formData = new FormData();
            formData.append('photo', this.photo);
            formData.append('caption', this.caption);
            console.log(this.photo)
            console.log(this.caption)

            try {
                const response = await this.$axios.post(`/photos`, formData, {
                    headers: {
                            'Content-Type': 'multipart/form-data',
                            'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
                            },
                });
                    console.log('Photo uploaded successfully', response.data);
                    this.photoUploadResult = "Photo uploaded!";
                    this.uploadSuccess = true;
                    this.buttonText = "Photo Uploaded!";
                    this.disabled = true;
                   setTimeout(() => {
                    this.buttonText = "Upload";
                    this.disabled = false;
                    this.$router.push('/users/' + sessionStorage.getItem('username'));
                }, 2000); 
                
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
    <div class="uploadForm">
        <div>
        <h2>Upload Photos</h2>
        <form class='uploadPhotoForm' @submit.prevent="upload"  >
          <div class="form-group">
            <label for="photo">Choose Photo:</label>
            <input type="file" id="photo"  @input="updatePhotoValue"  accept="image/*" required>
          </div>
          <div class="form-group">
            <label for="caption">Caption:</label>
            <input type="text" @input="updateCaptionValue" id="caption" >
          </div>
          <button class="uploadButton" :disabled="disabled"  type="submit">{{buttonText}}</button>
        </form>
    </div>
    </div>
</template>
<style>
.uploadForm{
   margin: 10% auto;
   align-items: center;
   align-self: center;
   justify-content: center;
   display: flex;
    height: 100%;
}
.uploadPhotoForm {
  width: 300px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f9f9f9;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input[type="file"],
#caption[type="text"],
.uploadButton {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.uploadButton {
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}

.uploadButton:hover {
  background-color: #45a049;
}
</style>