import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import NewPost from '../views/NewPost.vue'
import StreamView from '../views/StreamView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		
		{path: '/', component: LoginView},
		{path: '/users/:username', component: ProfileView},
		{path: '/photos', component: NewPost},
		{path: '/stream', component: StreamView},
	
	]
})

export default router
