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
              <span>Hashtag</span>
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

          <v-row justify="center" v-show="searchUsers">
            <v-list>
              <v-list-item v-for="user in allUsers" :key="user.Username">
                <v-card height="60" width="600" class="ma-3 grey lighten-5">
                  <v-card-title height="60" class="grey lighten-3">
                    <v-row>
                      <v-col>
                        <v-row>
                          <h4>@{{ user.Username }}</h4>
                        </v-row>
                        <v-row>
                          <label>{{user.Name}} {{space}} {{user.Surname}}</label>
                        </v-row>
                      </v-col>
                      <v-spacer/>
                      <v-col>
                        <v-btn small @click="loadUserProfile(user.Username)">
                          <span>View profile</span>
                        </v-btn>
                      </v-col>
                    </v-row>
                  </v-card-title>
                  <v-divider></v-divider>
                </v-card>
              </v-list-item>
            </v-list>
          </v-row>

          <v-row justify="center">
            <v-list>
              <v-list-item v-for="post in allPublicPosts" :key="post.Username">
                <v-card height="750" width="550" class="ma-3 grey lighten-5">
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
                  <v-row  class="ma-2" height="5">
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
                      >Hashtags</v-btn
                    >
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="viewAllTags(post.Tags)"
                      >Tags</v-btn
                    >
                    <v-btn
                      x-small
                      class="mr-3"
                      @click="commentPost(post.Id)"
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

            <v-btn class= "mx-2" @click="$router.go()">
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
  </v-app>
</template>

<script>
import { getId, getToken, getUsername } from "../security/token.js";
import AllPostComments from "../components/AllPostComments.vue";
import AddPostComment from "../components/AddPostComment.vue";
import AllTags from "../components/AllTags.vue";
import axios from "axios";
import AllHashtags from "../components/AllHashtags.vue";


export default {
  name: "Search",
  components: { AllPostComments, AddPostComment, AllTags, AllHashtags },

  data() {
    return {
      loggedUser: false,
      allPublicPosts: [],
      allUsers: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allHashtagsDialog: false,
      allPostComments: [],
      allPostTags: [],
      allPostHashtags: [],
      allTagsDialog: false,
      showUser: true,
      showTag: false,
      postId: 0,
      showLocation: false,
      searchInput: "",
      space: " ",
      showPublicPosts: true
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
    loadUserProfile(username){
      this.$router.push('/user-profile/' + username).catch(()=>{})
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
    searchUsers() {
      this.showUser = true
      this.showTag = false
      this.showLocation = false
    },
    searchTags() {
      this.showUser = false
      this.showTag = true
      this.showLocation = false
    },
    searchLocations() {
      this.showUser = false
      this.showTag = false
      this.showLocation = true
    },
    search() {
      if (this.showUser === true){
        axios.get("http://localhost:8081/api/user/search-public-regular-users/"+this.searchInput, {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allUsers = response.data;
          this.allPublicPosts = []
        })
        .catch((error) => {
          if (error.response.status === 500) {
            console.log("Internal server error");
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
          this.allUsers = []      
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
          this.allUsers = []      
        }).catch((error) => {
          if (error.response.status === 500) {
            console.lof("Internal server error");
          }
        });
        
      }
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
    this.loadAllPublicPosts();
  },
};
</script>

<style>
</style>