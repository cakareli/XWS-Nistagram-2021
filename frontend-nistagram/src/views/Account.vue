<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-card>
            <v-toolbar height="45" width="800px" >
              <v-app-bar  app height="45" color="grey">
                <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
                <v-row >
                  <v-col>
                    <v-toolbar-title>
                      <span outlined class="font-weight-light">NISTA</span>
                      <span>GRAM</span>
                    </v-toolbar-title>
                  </v-col>
                  <v-spacer></v-spacer>
                  <v-col>
                    <v-btn outlined class="mx-5 white" small>
                      <v-icon>mdi-send</v-icon>
                    </v-btn>
                  </v-col>
                </v-row>
              </v-app-bar>
              <v-navigation-drawer  v-model="drawer" absolute temporary height="570" class="grey lighten-4">
                <v-list>
                  <v-list-item-group
                    active-class="deep-purple--text text--accent-4">
                    <v-list-item>
                      <v-list-item-title>
                        <v-icon small>mdi-clock</v-icon>
                        Activity
                      </v-list-item-title>
                    </v-list-item>

                    <v-list-item @click="$router.push('/update')">
                      <v-list-item-title>
                        <v-icon small>mdi-pencil</v-icon>
                        Edit Profile
                        </v-list-item-title>
                    </v-list-item>

                    <v-list-item>
                      <v-list-item-title>
                        <v-icon small>mdi-lock</v-icon>
                        Privacy</v-list-item-title>
                    </v-list-item>

                    <v-list-item>
                      <v-list-item-title>
                        <v-icon small>mdi-bookmark</v-icon>
                        Saved Posts</v-list-item-title>
                    </v-list-item>

                    <v-list-item>
                      <v-list-item-title>
                        <v-icon small>mdi-text-box-multiple</v-icon>
                        Verify</v-list-item-title>
                    </v-list-item>

                    <v-list-item @click="logout">
                      <v-list-item-title>
                        <v-icon small>mdi-logout</v-icon>
                        Log out</v-list-item-title>
                    </v-list-item>

                  </v-list-item-group>
                </v-list>
              </v-navigation-drawer>
            </v-toolbar>
          </v-card>
            <v-container height="530" class="grey lighten-2">
              <v-row>
                <v-col width="400px"></v-col>
                <v-col>
                  <h3>{{this.username}}</h3>
                </v-col>
              </v-row>
              <v-row>
                <v-col width="400px"></v-col>
                <v-col>
                  <h3>{{this.name}}{{this.surname}}</h3>
                </v-col>
              </v-row>
                  
          </v-container>

          <v-card class="grey lighten-2" height="530">
            <v-virtual-scroll height="530" item-height="500" item-width="500" :items="allUserPosts">
              <template v-slot:default="{ item }" >              
                    <v-card class="grey lighten-2">
                      <v-card-title>
                        <h4>{{item.username}}</h4>
                      </v-card-title>
                      <v-row class="ma-2" justify="center">
                          <v-img v-bind:src="item.MediaPaths[0]" max-width="400" max-height="400"></v-img>
                      </v-row>
                      <v-row>
                        <v-col width="200"></v-col>
                        <v-col width="400">
                          <v-textarea label="Description" outlined v-model="item.Description" readonly ></v-textarea>
                        </v-col>
                        <v-col width="200"></v-col>                       
                      </v-row>
                      <v-row class="ma-2">
                        <span> Likes: {{item.Likes}}</span>
                        <v-spacer/>
                        <span> Dislikes: {{item.Dislikes}}</span>
                      </v-row>
                      
                      <v-row class="ma-2">
                        <v-btn class="mr-3" @click="likePost(item.Id)">Like</v-btn>
                        <v-btn @click="dislikePost(item.Id)" >Dislike</v-btn>
                        <v-spacer/>
                        <v-btn class="mr-3" @click="viewAllTags(item.Id)" >Tags</v-btn>
                        <v-btn class="mr-3" @click="commentPost(item.Id)" >Comment</v-btn>
                        <v-btn @click="viewAllPostComments(item.Id)">View all comments</v-btn>
                      </v-row>
                    </v-card>                
                <AllPostComments :allPostCommentsDialog.sync="allPostCommentsDialog"/>
                <AddPostComment :addPostCommentDialog.sync="addPostCommentDialog" /> 
                <AllTags :allTagsDialog.sync="allTagsDialog" />
                <v-divider></v-divider>
              </template>             
            </v-virtual-scroll>
          </v-card>

          <v-bottom-navigation height="35" width="800px" background-color="grey">
            <v-btn value="home" @click="$router.push('/')">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search" @click="$router.push('/search')">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn value="add"  @click="$router.push('/new-post')">
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn value="notification">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            
            <v-btn value="profile">
              <v-icon>mdi-account</v-icon>
            </v-btn>
              
          </v-bottom-navigation>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>

import axios from "axios";
import { getId, getToken, getUsername, removeToken } from '../security/token.js'
import AllPostComments from '../components/AllPostComments.vue'
import AddPostComment from '../components/AddPostComment.vue'
import AllTags from '../components/AllTags.vue'

export default {
  name: "Account",
  components:
        {AllPostComments, AddPostComment, AllTags},
  data() {
    return{
      drawer : false,
      name: "",
      surname: "",
      username: "",
      biography: "",
      website: "",
      allUserPosts: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allTagsDialog: false,
    }
  },
  methods: {
    loadRegisteredUser(){
            axios.get('http://localhost:8081/api/user/logged-user/'+getId(),
            {
                    headers : {
                        Authorization: 'Bearer ' + getToken()
                    }
            })
            .then(response => {
                this.regularUser = response.data;
                this.name = this.regularUser.Name + " ";
                this.surname = this.regularUser.Surname;
                this.username = this.regularUser.Username;
                this.biography = this.regularUser.Biography;
                this.website = this.regularUser.WebSite;
            }).catch(error => {
            if(error.response.status === 500){
                this.snackbarText = "Internal server error occurred!";
                this.snackbar = true;
            }
            if(error.response.status === 401){
               this.snackbar = true
               this.snackbarText = "You are unauthorized to get patient informations!";
            }
        })
    },
    logout(){
      removeToken()
      this.$router.push("/").catch(()=>{})
    },
    loadUsersPosts(){
      axios.get('http://localhost:8081/api/media-content/regular-user-posts/'+getUsername(),
      {
            headers : {
                          Authorization: 'Bearer ' + getToken()
                      }
      }).then(response => {
        this.allUserPosts = response.data;
        alert(this.allUserPosts.MediaPaths[0])

      }).catch(error => { 
        if(error.response.status === 500){
        console.lof("Internal server error")
        }
    })
    },
    likePost(postId) {
      console.log(postId)
    },
    dislikePost(postId){
      console.log(postId)
    },
    commentPost(postId) {
      this.postId = postId
      this.addPostCommentDialog = true;
    },
    viewAllPostComments(postId){
      console.log(postId)
      this.allPostCommentsDialog = true;
    },
    viewAllTags(postId){
      console.log(postId)
      this.allTagsDialog = true;
    },
  },
  mounted(){
    this.loadRegisteredUser();
    this.loadUsersPosts();
  }
};
</script>

<style>
</style>