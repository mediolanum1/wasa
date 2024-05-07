
<script>
export default {
		data(){
			return {	
				logging: false,
                username: "",
                loginStatus: "Login",
			}
        
		}, 
        async mounted(){
            
        },

        methods: { 
          
         async login() {
            this.logging = true;
            this.error = null;
            try {
                let response = await this.$axios.post('/sessions', { username: this.username }, {
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                    },
                });
            
                sessionStorage.setItem('token', response.data)
                sessionStorage.setItem('username', this.username)
                sessionStorage.setItem('loggedStatus', true)
                this.loginStatus = "You're Logged-in"

                    this.logging = false;
                    localStorage.setItem('loginExecuted', '1')
                    location.reload();
                    this.$router.push('/users/' + this.username);
             
                
                
            } catch (error) {
                this.loginStatus = "Error Logging-In"
                console.log("Error while logging"); 
                setTimeout(() => {
                    this.loginStatus = "Login";
                    this.logging = false;
                }, 2000);    
            }

            }
        }
	}

	
</script>

<template>
    <div class='login-view'> 
    <div class="login-container">
        <h2>Login</h2>
        <form action="#" method="post">
            <div class="form-group">
                <label for="username">Username</label>
                <input v-model="username" type="text" id="username" required>
            </div>
           
            <div class="form-group">
                <button type="submit" @click="login" :disabled="logging">{{ loginStatus }}</button> 
            </div>
        </form>
    </div>
</div>
</template>

<style>
.login-view{
    width: 100%;
    height: 100%;
}
.login-container {
    margin:  15% auto;  
    width: 300px;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
.login-container h2 {
    text-align: center;
    margin-bottom: 20px;
}
.form-group {
    margin-bottom: 15px;
}
.form-group label {
    display: block;
    margin-bottom: 5px;
}
.form-group input, button{
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 3px;
    height: 7vh;
}
.form-group button{
    background-color: #007bff;
    color: #fff;
    cursor: pointer;
}
.form-group button:hover {
    background-color: #0056b3;
}

</style>