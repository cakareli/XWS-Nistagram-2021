<template>
  <v-app class="grey lighten-2">
    <v-app-bar app height="45" class="grey lighten-3 justify-center">
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
          <v-bottom-navigation background-color="grey lighten-3" height="45px">
            <v-btn @click="searchUsers">
              <v-icon>mdi-account</v-icon>
              <span>User</span>
            </v-btn>

            <v-btn @click="searchTags">
              <v-icon>mdi-pound-box</v-icon>
              <span>Tag</span>
            </v-btn>

            <v-btn @click="searchLocations">
              <v-icon>mdi-map-marker</v-icon>
              <span>Location</span>
            </v-btn>
          </v-bottom-navigation>
          <v-container height="30px">
            <v-row height="30px">
              <v-text-field
                prepend-icon="mdi-account"
                v-model="searchInput"
                placeholder="Search users"
                v-show="showUser"
              ></v-text-field>
              <v-text-field 
                prepend-icon="mdi-pound-box"
                v-model="searchInput"
                placeholder="Search tags"
                v-show="showTag">
              </v-text-field>
              <v-text-field
                prepend-icon="mdi-map-marker"
                v-model="searchInput"
                placeholder="Search locations"
                v-show="showLocation"
              ></v-text-field>
              <v-btn height="30" class="ma-5" @click="search">
                <v-icon>mdi-magnify</v-icon>
              </v-btn>
            </v-row>
          </v-container>
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="post in allPublicPosts" :key="post.Username">
                <v-card height="665" width="500" class="ma-3 grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                    <h4>@{{ post.RegularUser.Username }}</h4>
                  </v-card-title>
                  <v-icon class="ml-11">mdi-map-marker</v-icon> {{post.Location}}
                  <v-row class="justify-center my-1">
                    <v-img
                      v-bind:src="post.MediaPaths[0]"
                      max-width="400"
                      max-height="400"
                      width="400"
                      height="400"
                      class="ma-2"
                    ></v-img>
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
            <v-btn class= "mx-2" @click="$router.push('/').catch(()=>{})">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn class= "mx-2">
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
  </v-app>
</template>

<script>
import { getId, getToken } from "../security/token.js";
import AllPostComments from "../components/AllPostComments.vue";
import AddPostComment from "../components/AddPostComment.vue";
import AllTags from "../components/AllTags.vue";
import axios from "axios";

export default {
  name: "Search",
  components: { AllPostComments, AddPostComment, AllTags },

  data() {
    return {
      loggedUser: false,
      allPublicPosts: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allPostComments: [],
      allPostTags: [],
      allTagsDialog: false,
      showUser: true,
      showTag: false,
      postId: 0,
      showLocation: false,
      searchInput: "",
    };
  },

  methods: {
    checkLoggedUser() {
      if (getId().length != 0) {
        this.loggedUser = true;
      }
    },
    loadAllPublicPosts() {
      axios.get("http://localhost:8081/api/media-content/public-posts", {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allPublicPosts = response.data;
        });
    },
    likePost(postId) {
      console.log(postId);
    },
    dislikePost(postId) {
      console.log(postId);
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
    searchUsers() {
      (this.showUser = true),
        (this.showTag = false),
        (this.showLocation = false);
      this.searchInput = "";
    },
    searchTags() {
      (this.showUser = false),
        (this.showTag = true),
        (this.showLocation = false);
      this.searchInput = "";
    },
    searchLocations() {
      (this.showUser = false),
        (this.showTag = false),
        (this.showLocation = true);
      this.searchInput = "";
    },
    search() {
      if (this.showUser === true){
        axios.get("http://localhost:8081/api/media-content/search-user/"+this.searchInput, {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allPublicPosts = response.data;
        })
        .catch((error) => {
          if (error.response.status === 500) {
            console.lof("Internal server error");
          }
        });
        
      }
      else if(this.showTag === true){
        axios.get("http://localhost:8081/api/media-content/search-tag/"+this.searchInput, {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allPublicPosts = response.data;
        })
        .catch((error) => {
          if (error.response.status === 500) {
            console.lof("Internal server error");
          }
        });
        
      }else if(this.showLocation === true){
        axios.get("http://localhost:8081/api/media-content/search-location/"+this.searchInput, {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allPublicPosts = response.data;
        }).catch((error) => {
          if (error.response.status === 500) {
            console.lof("Internal server error");
          }
        });
        
      }
    },
  },
  mounted() {
    this.checkLoggedUser();
    this.loadAllPublicPosts();
  },
};
</script>

<style>
</style>