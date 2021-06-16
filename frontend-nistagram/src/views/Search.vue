<template>
    <v-app class="grey lighten-2">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45"  width="800px" >
            <v-app-bar app height="45" color="grey">
              <v-row>
                <v-col>
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
                <v-spacer></v-spacer>
                <v-col>
                  <v-btn outlined class="mx-5 white" v-show="loggedUser">
                    <v-icon>mdi-send</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-toolbar height="40">
            <v-btn @click="searchUsers">
                <v-icon>mdi-account</v-icon>
                <span>User</span>               
            </v-btn>

            <v-btn  @click="searchTags">
                <v-icon>mdi-tag</v-icon>
                <span>Tag</span>               
            </v-btn>

            <v-btn @click="searchLocations">
                <v-icon>mdi-map</v-icon>
                <span>Location</span>               
            </v-btn>
          </v-toolbar>
          <v-container height="30px">
            <v-row height="30px">
                <v-text-field v-model="searchInput" placeholder="Search users" v-show="showUser"></v-text-field>
                <v-text-field v-model="searchInput" placeholder="Search tags" v-show="showTag"></v-text-field>
                <v-text-field v-model="searchInput" placeholder="Search locations" v-show="showLocation"></v-text-field>
                <v-btn height="30" class="mt-5" @click="search">
                    <v-icon>mdi-magnify</v-icon>
                </v-btn>
            </v-row>
          </v-container>
          <v-card class="grey lighten-2" height="530">
            <v-virtual-scroll height="530" item-height="500" item-width="500" :items="allPublicPosts" v-show="loggedUser">
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
                        <v-btn class="mr-3" @click="likePost(item.Id)" v-show="loggedUser">Like</v-btn>
                        <v-btn @click="dislikePost(item.Id)" v-show="loggedUser">Dislike</v-btn>
                        <v-spacer/>
                        <v-btn class="mr-3" @click="viewAllTags(item.Id)" v-show="loggedUser">Tags</v-btn>
                        <v-btn class="mr-3" @click="commentPost(item.Id)" v-show="loggedUser">Comment</v-btn>
                        <v-btn @click="viewAllPostComments(item.Id)">View all comments</v-btn>
                      </v-row>
                    </v-card>                
                <AllPostComments :allPostCommentsDialog.sync="allPostCommentsDialog"/>
                <AddPostComment :addPostCommentDialog.sync="addPostCommentDialog" :postId = "postId"/> 
                <AllTags :allTagsDialog.sync="allTagsDialog" :postId = "postId"/>
                <v-divider></v-divider>
              </template>             
            </v-virtual-scroll>
          </v-card>
          <v-bottom-navigation height="35" width="800px" background-color="grey">
            <v-btn value="home" @click="$router.push('/')">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn value="add" @click="$router.push('/new-post')" v-show="loggedUser">
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn value="notification" v-show="loggedUser">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            
            <v-btn
              v-show="loggedUser"
              value="profile"
              @click="$router.push('/account')"
            >
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

import {getId, getToken} from '../security/token.js'
import AllPostComments from '../components/AllPostComments.vue'
import AddPostComment from '../components/AddPostComment.vue'
import AllTags from '../components/AllTags.vue'
import axios from "axios";

export default {
    name: 'Search',
    components:
        {AllPostComments, AddPostComment, AllTags},

    data() {
    return {
      loggedUser: false,
      allPublicPosts : [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allTagsDialog: false,
      showUser: true,
      showTag: false,
      showLocation: false,
      searchInput: "",
    }
  },

  methods: {
    checkLoggedUser(){
      if(getId().length != 0){
        this.loggedUser = true
      }      
    },
    loadAllPublicPosts() {
        axios.get('http://localhost:8081/api/media-content/public-posts',{
            headers : {
                        Authorization: 'Bearer ' + getToken()
                    }
        })
        .then(response => {
            
            this.allPublicPosts = response.data
                     
            
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
    searchUsers(){
        this.showUser= true,
        this.showTag= false,
        this.showLocation= false
        this.searchInput = ""
    },
    searchTags(){
        this.showUser= false,
        this.showTag= true,
        this.showLocation= false
        this.searchInput = ""
    },
    searchLocations(){
        this.showUser= false,
        this.showTag= false,
        this.showLocation= true
        this.searchInput = ""
    },
    search(){
    },

  },
  mounted() {
    this.checkLoggedUser();
    this.loadAllPublicPosts();
  },
}
</script>

<style>

</style>