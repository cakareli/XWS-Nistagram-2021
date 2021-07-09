<template>
    <v-app class="grey lighten-2" width="800px">
      <v-bottom-navigation background-color="grey lighten-3" height="45px">
            <v-btn @click="$router.push('/notifications')" class="grey lighten-2" left>Go Back</v-btn>
        </v-bottom-navigation>
        <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="post in postFromNotification" :key="post.id">
                <v-card height="790" width="550" class="grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                    <a :href="'/user-profile/' + post.RegularUser.Username" class="black--text" style="text-decoration: none; color: inherit;">@{{ post.RegularUser.Username }}</a>
                    <v-spacer/>
                    <v-btn small  @click="savePost(post.Id)">
                      <v-icon>mdi-bookmark</v-icon>
                    </v-btn>
                  </v-card-title>
                  <v-icon class="ml-11">mdi-map-marker</v-icon> {{post.Location}}
                  <v-row class="justify-center my-1" v-show="post.MediaContentType==0">
                    <v-img
                      v-bind:src="post.MediaPaths[0]"
                      max-width="400"
                      max-height="400"
                      width="400"
                      height="400"
                      class="ma-2"
                    ></v-img>
                  </v-row>
                  <v-row class="justify-center my-1" v-show="post.MediaContentType==2">
                    <v-carousel class="mx-16" >
                      <v-carousel-item
                      height="300px"
                      v-for="image in post.MediaPaths"
                      :key="image"
                      :src="image">                     
                      </v-carousel-item>
                    </v-carousel>
                  </v-row>
                  <v-row class="justify-center mx-2">
                    <v-textarea
                      label="Description"
                      v-model="post.Description"
                      auto-grow
                      outlined
                      rows="1"
                      width="170"
                      readonly
                    ></v-textarea>
                  </v-row>
                  <v-row class="ma-2" height="5">
                    <span> Likes: {{ post.Likes }}</span>
                    <v-spacer />
                    <span> Dislikes: {{ post.Dislikes }}</span>
                  </v-row>
                  <v-row class="ma-2">
                    <v-btn x-small class="mr-3" @click="likePost(post.Id)" v-show="loggedUser">
                      <v-icon x-small left>mdi-thumb-up</v-icon>
                      <span>Like</span>
                    </v-btn>

                    <v-spacer></v-spacer>
                    <v-spacer></v-spacer>

                    <v-btn
                      x-small
                      class="error"
                      @click="reportPost(post)"
                      v-show="loggedUser"
                    ><v-icon x-small left color="white">mdi-alert-octagon</v-icon>
                      <span>Report</span>
                      </v-btn>
                    
                    <v-spacer></v-spacer>
                    <v-spacer></v-spacer>

                    <v-btn
                      x-small
                      class="mr-3"
                      @click="dislikePost(post.Id)"
                      v-show="loggedUser"
                      >
                      <v-icon x-small left>mdi-thumb-down</v-icon>
                      <span>Dislike</span>
                      </v-btn >
                    
                  </v-row>
                  <br>
                  <v-row class="ma-2">
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllHashtags(post.Hashtags)"
                      >Hashtags</v-btn
                    >
                    <v-spacer></v-spacer>
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllTags(post.Tags)"
                      >Tags</v-btn
                    >
                    <v-spacer></v-spacer>
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="commentPost(post.Id)"
                      v-show="loggedUser"
                      >Comment</v-btn
                    >
                    <v-spacer></v-spacer>
                    <v-btn x-small @click="viewAllPostComments(post.Comment)"
                      >View all comments</v-btn
                    >
                  </v-row>
                </v-card>

                <v-divider></v-divider>
              </v-list-item>
            </v-list>
          </v-row>
        </v-card>
      </v-row>
    
    </v-container>


    
    <AllPostComments :allPostCommentsDialog.sync="allPostCommentsDialog" :allPostComments="allPostComments"/>
    <AddPostComment
      :addPostCommentDialog.sync="addPostCommentDialog"
      :postId="postId"
    />
    <AllTags :allTagsDialog.sync="allTagsDialog" :allPostTags="allPostTags"/>
    <AllHashtags :allHashtagsDialog.sync="allHashtagsDialog" :allPostHashtags="allPostHashtags"/>
    <ReportPost :reportPostDialog.sync="reportPostDialog" :reportedPost="reportedPost"/>
    <v-main>
      <router-view />
    </v-main>
    </v-app>
</template>

<script>

import axios from "axios";
import {getToken, getUsername } from "../security/token.js"
import AllPostComments from "../components/AllPostComments.vue"
import AllTags from "../components/AllTags.vue"
import AllHashtags from "../components/AllHashtags.vue"
import ReportPost from "../components/ReportPost.vue"
import AddPostComment from "../components/AddPostComment.vue"

export default{
    name: "ViewPostFromNotification",
    components: {
        AllPostComments,
        AllTags,
        AllHashtags,
        ReportPost,
        AddPostComment,
    },
    data() {
        return {
            loggedUser: false,
            postFromNotification: [],
            postFromNotificationId: 0,
            allPostCommentsDialog: false,
            addPostCommentDialog: false,
            allTagsDialog: false,
            allHashtagsDialog: false,
            viewStoryDialog: false,
            allPostTags: [],
            allPostHashtags: [],
            allPostComments: [],
            storyView: {},
            reportedPost: {},
            reportPostDialog: false,
            postId: ""
        }
    },
    created() {
        this.postFromNotificationId = this.$route.params.post;
    },
    mounted(){
        this.loadPostFromNotification()
    },
    methods: {
        loadPostFromNotification(){
            alert(this.postFromNotificationId)
        },
        reportPost(post){
            this.reportPostDialog = true;
            this.reportedPost = post;
        },
        likePost(postId) {
            console.log(postId);
            let likePostDTO = {
            username: getUsername(),
            postId: postId,
        }
        axios.put("http://localhost:8081/api/media-content/like-post",
            likePostDTO,
        {
          headers: {
            Authorization: "Bearer " + getToken(),
        },
        })
        .then((response) => {
            console.log(response)
            this.$router.go()
        });
    },
    dislikePost(postId) {
      console.log(postId);
      let dislikePostDTO = {
        username: getUsername(),
        postId: postId,
      }
        axios.put("http://localhost:8081/api/media-content/dislike-post",
            dislikePostDTO,
        {
          headers: {
            Authorization: "Bearer " + getToken(),
        },
        })
        .then((response) => {
            console.log(response)
            this.$router.go()
        });
    },
    commentPost(postId) {
      this.postId = postId;
      this.addPostCommentDialog = true;
    },
    viewAllPostComments(allPostComments) {
      this.allPostComments = allPostComments;
      this.allPostCommentsDialog = true;
    },
    viewAllTags(allPostTags) {
      this.allPostTags = allPostTags;
      this.allTagsDialog = true;
    },
    viewAllHashtags(allPostHashtags) {
      this.allPostHashtags = allPostHashtags;
      this.allHashtagsDialog = true;
    },
    savePost(postId){
      console.log(postId);
    },
    }
}
</script>
