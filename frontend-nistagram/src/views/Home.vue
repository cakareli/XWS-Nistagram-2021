<template>
  <v-app class="grey lighten-2" width="800px">
    <v-app-bar app height="45" class="grey lighten-3">
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

    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">
          <v-bottom-navigation background-color="grey lighten-3" height="45px" v-show="!loggedUser" >
            <v-btn @click="$router.push('/registration')" class="grey lighten-2">Register</v-btn>
            <v-spacer></v-spacer>
              <v-btn @click="$router.push('/login')" class="grey lighten-2">Login</v-btn>
          </v-bottom-navigation>
          <v-row justify="center">
            <v-card height="160px" width="550px" class="pa-5 grey lighten-5">
                <v-row height="160px">
                  <v-slide-group multiple show-arrows class="mt-9" >
                    <v-slide-item  v-for="(story, index) in allStories" :key="story.Id">
                      <v-btn height="80" fab icon class="mx-5" @click="viewStory(index)">
                        <v-avatar size="70px">
                          <v-img v-bind:src="story.MediaPaths[0]"></v-img>
                        </v-avatar>
                      </v-btn>
                    </v-slide-item>
                  </v-slide-group>
                </v-row>
            </v-card>
          </v-row>
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="post in allPublicPosts" :key="post.Id">
                <v-card height="750" width="550" class="grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                    <h4>@{{ post.RegularUser.Username }}</h4>
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
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="likePost(post.Id)"
                      v-show="loggedUser"
                      >Like</v-btn
                    >
                    <v-btn
                      x-small
                      @click="dislikePost(post.Id)"
                      v-show="loggedUser"
                      >Dislike</v-btn
                    >
                    <v-spacer />
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllHashtags(post.Hashtags)"
                      v-show="loggedUser"
                      >Hashtags</v-btn
                    >
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllTags(post.Tags)"
                      v-show="loggedUser"
                      >Tags</v-btn
                    >
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="commentPost(post.Id)"
                      v-show="loggedUser"
                      >Comment</v-btn
                    >
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
    
    <v-footer app height="45px" class="grey lighten-3 justify-center">
        <v-container>
          <v-row justify="center">
            <v-btn class= "mx-2" @click="$router.go()">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search" class= "mx-2" @click="$router.push('/search').catch(()=>{})">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2"
              @click="$router.push('/new-post').catch(()=>{})"
              v-show="loggedUser"
            >
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn class= "mx-2" v-show="loggedUser">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            <v-btn
              v-show="loggedUser"
              class= "mx-2"
              @click="$router.push('/account').catch(()=>{})"
            >
              <v-icon>mdi-account</v-icon>
            </v-btn>
          </v-row>
        </v-container>
    </v-footer>
    <AllPostComments :allPostCommentsDialog.sync="allPostCommentsDialog" :allPostComments="allPostComments"/>
    <AddPostComment
      :addPostCommentDialog.sync="addPostCommentDialog"
      :postId="postId"
    />
    <AllTags :allTagsDialog.sync="allTagsDialog" :allPostTags="allPostTags"/>
    <AllHashtags :allHashtagsDialog.sync="allHashtagsDialog" :allPostHashtags="allPostHashtags"/>
    <ViewStory :viewStoryDialog.sync="viewStoryDialog" :storyView="storyView" />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>

import axios from "axios";
import { getId, getToken, getUsername} from "../security/token.js";
import AllPostComments from "../components/AllPostComments.vue";
import AddPostComment from "../components/AddPostComment.vue";
import AllTags from "../components/AllTags.vue";
import AllHashtags from "../components/AllHashtags.vue";
import ViewStory from "../components/ViewStory.vue";


export default {
  name: "Home",
  components: {
    AllPostComments,
    AddPostComment,
    AllTags,
    AllHashtags,
    ViewStory
  },
  data() {
    return {
      loggedUser: false,
      allStories: [],
      allPublicPosts: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allTagsDialog: false,
      allHashtagsDialog: false,
      viewStoryDialog: false,
      allPostTags: [],
      allPostHashtags: [],
      allPostComments: [],
      storyView: {},
      postId: ""
    }
  },

  methods: {

    viewStory(story){
      this.storyView = this.allStories[story]
      this.viewStoryDialog = true
    },

    loadAllPublicPostsForGuest() {
      axios
        .get("http://localhost:8081/api/media-content/public-posts", {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allPublicPosts = response.data;
        });
    },
    loadAllStories(){
      axios
        .get("http://localhost:8081/api/media-content/regular-user-stories/"+getUsername(), {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allStories = response.data;
        }).catch(error => {
            if(error.response.status === 500){
                this.snackbarText = "Internal server error occurred!";
                this.snackbar = true;
            }
        })
    },
    
    checkLoggedUser() {
      if (getId().length != 0) {
        this.loggedUser = true;
      }
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
      let savePostDTO = {
        username: getUsername(),
        postId: postId,
        isAdd: "yes",
      }
        axios.put("http://localhost:8081/api/user/save-post",
            savePostDTO,
        {
          headers: {
            Authorization: "Bearer " + getToken(),
        },
        })
        .then((response) => {
            console.log(response)
            this.$router.go()
        });
    }
  },
  mounted() {
    this.checkLoggedUser();
    this.loadAllPublicPostsForGuest();
    this.loadAllStories();
  },
};
</script>
