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
                  <v-btn outlined class="mx-5 white">
                    <v-icon>mdi-send</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-toolbar id="guestToolbar" ref="guestToolbar" height="35" class="grey lighten-4" width="800px"  v-show="!loggedUser">
            <v-app-bar app>
              <v-row align="center" justify="space-around">
                <v-col>
                  <v-btn 
                    width="250px"
                    height="35px"
                    @click="$router.push('/registration')"
                    class="grey lighten-2"
                    >Register</v-btn
                  >
                </v-col>
                <v-spacer></v-spacer>
                <v-col>
                  <v-btn
                    width="250px"
                    height="35px"
                    @click="$router.push('/login')"
                    class="grey lighten-2"
                    >Login</v-btn
                  >
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-row class="ma-3" allign="center" justify="center" v-for= "item in allPublicPosts" :key = "item.username">
                <v-card class= "ma-3">
                  <v-card-title>
                    <h4>{{item.username}}</h4>
                  </v-card-title>
                    <v-row class="ma-2">
                         <v-img src="//placehold.it/600x500" max-width="600"></v-img>
                    </v-row>
                    <v-row class="ma-2">
                      <span> Likes: {{item.likes}}</span>
                      <v-spacer/>
                      <span> Dislikes: {{item.dislikes}}</span>
                    </v-row>
                    <v-row class="ma-2">
                      <v-btn class="mr-3" @click="likePost(item.id)">Like</v-btn>
                      <v-btn @click="dislikePost(item.id)">Dislike</v-btn>
                      <v-spacer/>
                      <v-btn class="mr-3" @click="commentPost(item.id)">Comment</v-btn>
                      <v-btn @click="viewAllPostComments(item.id)">View all comments</v-btn>
                    </v-row>
                </v-card>
            </v-row>
            <AllPostComments :allPostCommentsDialog.sync="allPostCommentsDialog"/>
            <AddPostComment :addPostCommentDialog.sync="addPostCommentDialog" :postId = "postId"/>
          </v-container>
          <v-bottom-navigation height="35" width="800px" background-color="grey">
            <v-btn value="home">
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
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>

import { getId } from '../security/token.js'
import AllPostComments from '../components/AllPostComments.vue'
import AddPostComment from '../components/AddPostComment.vue'

export default {
  name: "Home",
  components: {
    AllPostComments,
    AddPostComment,
  },
  data() {
    return {
      loggedUser: false,
      allPublicPosts : [
        {id: "1", username: "denisfruza", likes: 21, dislikes: 11, imagepath:"aaaaaaaaaaaaaa"},
        {id: "2", username: "dbulaja98", likes: 22, dislikes: 12, imagepath:"bbbbbbbbbbbbbbbbbb"},
        {id: "3", username: "cakinjoo", likes: 23, dislikes: 13, imagepath:"cccccccccccccccccccc"},
        {id: "4", username: "jarulja", likes: 24, dislikes: 14, imagepath:"cccccccccccccccccccc"}
      ],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      postId: "",
    }
  },

  methods: {
    checkLoggedUser(){
      if(getId().length != 0){
        this.loggedUser = true
      }      
    },
    loadAllPublicPosts() {

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
  },
  mounted() {
    this.checkLoggedUser();
  },
};

</script>
