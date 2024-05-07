<script setup>
import { RouterLink, RouterView } from 'vue-router';
</script>


<script> 
    export default {
        data(){
            return {
                searchBar: false,
                searchUser: "",
            }
        },

        async mounted(){
            if (sessionStorage.getItem('token') != null) {
            this.searchBar = true;
            if (sessionStorage.getItem('loggedStatus')){
                sessionStorage.removeItem('loggedStatus')
                 
                if (sessionStorage.getItem('userSearch')== null){
                    
                    this.$router.push('/users/' + sessionStorage.getItem('username'));
                    }else{
                        this.$router.push('/users/' + sessionStorage.getItem('userSearch'));
                    }
            }
            }else{
                this.$router.push('/');
            }
        },
        methods: {
            logout(){
                    localStorage.clear();
                    sessionStorage.clear();
                    location.reload();
                },
                search() {
                    sessionStorage.setItem('userSearch',this.searchUser)
                    this.$router.push('/users/' + this.searchUser);
                },
                profile(){
                    sessionStorage.removeItem('userSearch')
                    this.$router.push('/users/' + sessionStorage.getItem('username'));
                },
                newPost(){
                    this.$router.push('/photos');
                },
                stream(){
                    this.$router.push('/stream');
                }
        }
    }
</script>


<template>
    <header v-if="searchBar" class="main-header">
        <div class="header-left">
            <!-- Icon (Replace 'icon.svg' with your icon) -->
            <svg class="icon">
            <use href="/feather-sprite-v4.29.0.svg#insta"/>
        </svg>
            <h3 @click="stream">WASAPhoto</h3>
        </div>
        <div class="header-center">
            <!-- Search bar -->
            <form  @submit.prevent="search">    <input type="text"  v-model="searchUser"  name='userSearchBar'  placeholder="Search..."><button class="searchButton" type="submit" @click="search" >Go</button></form>
        
            
            
        </div>
        <div class="header-right">
            <!-- Drop down menu -->
            <div class="dropdown">
                
                <button class="dropbtn">Menu</button>
                <div class="dropdown-content">
                    <p @click="profile" >Profile</p>
                    <p @click="newPost">New Post</p>
                    <p  @click="logout">Log-out</p>
                </div>
            </div>
        </div>
    </header>
        <body>
            <RouterView />
        </body>


</template>


<style>
html { height: 100%; overflow:auto; }
body{
    align-items: center;
    justify-content: center;
    width: 100vw;
    height: 100%;
}

.searchButton{
    width: 10%;
    
    border-radius: 5px;
    
}
* {
margin: 0;
padding: 0;
box-sizing: border-box;
}

.main-header {
display: flex;
justify-content: space-between;
align-items: center;
background-color: #333;
color: #fff;
padding: 10px 20px;
}

.header-left{
display: flex;
margin-left: 20px;
}
.header-left .icon{
width: 30px; 
height: 30px;
margin-top: 0 auto;
}
.header-left h3 {
margin-left: 10px;

}

.header-center input {
width: 300px; 
padding: 5px;
border-radius: 5px;
border: none;
outline: none;
margin: auto;
}

.header-right{
width: auto;
}

.header-right .dropdown {
position: relative;
display: inline-block;
margin-right: 30px;

}


.dropbtn {
background-color: transparent;
color: white;
border: 1px solid;
width: 100%;
cursor: pointer;
justify-content: center;

}
.dropdown{
margin-right:30px;
justify-content: center;
display: flex;

}

.dropdown-content {
display: none;
position: absolute;
background-color: #f9f9f9;
min-width: 100px;
z-index: 1;

border-radius: 5px;
}


.dropdown-content p {
color: black;
padding: 12px 16px;
text-decoration: none;

display: block;
}


.dropdown-content p:hover {
background-color: #ddd;
}

.dropdown:hover .dropdown-content {
display: block;
}

</style>