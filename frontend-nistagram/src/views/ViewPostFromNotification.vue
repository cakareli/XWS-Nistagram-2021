<template>
    <v-app class="grey lighten-2" width="800px">
      <v-app-bar app height="45" color="grey lighten-3">
              <v-app-bar-nav-icon @click="$router.push('/notifications')">
                <v-icon>mdi-arrow-left</v-icon>
              </v-app-bar-nav-icon>
              <v-row>
                <v-col>
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
              </v-row>
            </v-app-bar>
            <br/><br/>
      <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">
          <v-row justify="center">
            <v-list>
                <v-card height="790" width="550" class="grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                    <a :href="'/user-profile/' + postFromNotification.RegularUser.Username" class="black--text" style="text-decoration: none; color: inherit;">@{{ postFromNotification.RegularUser.Username }}</a>
                    <v-spacer/>
                  </v-card-title>
                  <v-icon class="ml-11">mdi-map-marker</v-icon> {{postFromNotification.Location}}
                  <v-row class="justify-center my-1" v-show="postFromNotification.MediaContentType==0 || postFromNotification.MediaContentType==1">
                    <v-img
                      v-bind:src="postFromNotification.MediaPaths[0]"
                      max-width="400"
                      max-height="400"
                      width="400"
                      height="400"
                      class="ma-2"
                    ></v-img>
                  </v-row>
                  <v-row class="justify-center my-1" v-show="postFromNotification.MediaContentType==2">
                    <v-carousel class="mx-16" >
                      <v-carousel-item
                      height="300px"
                      v-for="image in postFromNotification.MediaPaths"
                      :key="image"
                      :src="image">                     
                      </v-carousel-item>
                    </v-carousel>
                  </v-row>
                  <v-row class="justify-center mx-2">
                    <v-textarea
                      label="Description"
                      v-model="postFromNotification.Description"
                      auto-grow
                      outlined
                      rows="1"
                      width="170"
                      readonly
                    ></v-textarea>
                  </v-row>
                  <v-row class="ma-2" height="5" v-show="!story">
                    <span> Likes: {{ postFromNotification.Likes }}</span>
                    <v-spacer />
                    <span> Dislikes: {{ postFromNotification.Dislikes }}</span>
                  </v-row>
                  <br>
                  <v-row class="ma-2">
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllHashtags(postFromNotification.Hashtags)"
                      >Hashtags</v-btn
                    >
                    <v-spacer></v-spacer>
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllTags(postFromNotification.Tags)"
                      >Tags</v-btn
                    >                    
                    <v-spacer></v-spacer>
                    <v-btn x-small @click="viewAllPostComments(postFromNotification.Comment)"
                      >View all comments</v-btn
                    >
                  </v-row>
                </v-card>
            </v-list>
          </v-row>
        </v-card>
      </v-row>
    
    </v-container>


    
    <AllPostComments :allPostCommentsDialog.sync="allPostCommentsDialog" :allPostComments="allPostComments"/>
    <AllTags :allTagsDialog.sync="allTagsDialog" :allPostTags="allPostTags"/>
    <AllHashtags :allHashtagsDialog.sync="allHashtagsDialog" :allPostHashtags="allPostHashtags"/>
    <v-main>
      <router-view />
    </v-main>
    </v-app>
</template>

<script>

import axios from "axios";
import AllPostComments from "../components/AllPostComments.vue"
import AllTags from "../components/AllTags.vue"
import AllHashtags from "../components/AllHashtags.vue"

export default{
    name: "ViewPostFromNotification",
    components: {
        AllPostComments,
        AllTags,
        AllHashtags,
    },
    data() {
        return {
            loggedUser: false,
            postFromNotification: [],
            postFromNotificationId: 0,
            contentType: -1,
            allPostCommentsDialog: false,
            allTagsDialog: false,
            allHashtagsDialog: false,
            viewStoryDialog: false,
            allPostTags: [],
            allPostHashtags: [],
            allPostComments: [],
            storyView: {},
            postId: "",
            story: false,
        }
    },
    created() {
        this.postFromNotificationId = this.$route.params.post;
        this.contentType = this.$route.params.contentType;
    },
    mounted(){
        this.loadPostFromNotification()
    },
    methods: {
      loadPostFromNotification(){
        if(this.contentType != 1){
          axios.get("http://localhost:8081/api/media-content/find-post/"+this.postFromNotificationId,)
          .then((response)=>{
            this.postFromNotification = response.data
            console.log(response.data.status)
          })
          .catch((error)=>{
            if(error.status === 404){
              console.log("Status not found")
            }
          })
        }
        else if(this.contentType == 1){
          this.story = true
          axios.get("http://localhost:8081/api/media-content/find-story/"+this.postFromNotificationId,)
          .then((response)=>{
            this.postFromNotification = response.data
            console.log(response.data.status)
          })
          .catch((error)=>{
            if(error.status === 404){
              console.log("Status not found")
            }
          })
        }
          
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
    }
}
</script>
