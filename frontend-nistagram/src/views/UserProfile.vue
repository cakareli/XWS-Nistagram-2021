<template>
    <v-app class="grey lighten-2" width="800px">
        <v-app-bar app height="45" class="grey lighten-3">
        <v-app-bar-nav-icon @click="$router.push('/search')">
            <v-icon>mdi-arrow-left</v-icon>
        </v-app-bar-nav-icon>
        <v-row>
          <v-col>
            <v-toolbar-title>
              <span outlined class="font-weight-light">NISTA</span>
              <span>GRAM</span>
            </v-toolbar-title>
          </v-col>
          <v-spacer></v-spacer>
          <v-col>
            <v-btn outlined class="mx-15 white" v-show="loggedUser">
              <v-icon>mdi-send</v-icon>
            </v-btn>
          </v-col>
        </v-row>
    </v-app-bar>


    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12 grey lighten-4">
            <v-row justify="center">
                <v-col>
                    <h2>@{{this.username}}</h2>
                    <br>
                    <h3>{{this.name}} {{this.space}} {{this.surname}}</h3>
                    <br>
                    <a icon :href="`${this.webSite}`" target="_blank">
                        {{this.webSite}}
                    </a>
                    <br><br>
                    <v-textarea
                      label="Biography"
                      v-model="this.biography"
                      auto-grow
                      outlined
                      rows="1"
                      width="170"
                      readonly
                    ></v-textarea>
                </v-col>
            </v-row>
            <v-row justify="center">
                <v-btn class="mx-5">Follow</v-btn>
                <v-btn  class="mx-5">Unfollow</v-btn>
                <v-btn  class="mx-5">Mute</v-btn>
                <v-btn  class="mx-5">Block</v-btn>
            </v-row>
        </v-card>
      </v-row> 
    </v-container>


    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">
            <v-row justify="center" >
            <v-list>
              <v-list-item v-for="post in allUserPosts" :key="post.Username">
                <v-card height="665" width="500" class="ma-3 grey lighten-5">
                  <v-card-title class="grey lighten-3" height="10">
                    <h4>@{{ post.RegularUser.Username }}</h4>
                    <v-spacer/>
                    <v-btn small  @click="savePost(post.Id)">
                      <v-icon>mdi-bookmark</v-icon>
                    </v-btn>
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

            <v-btn class= "mx-2" @click="$router.push('/search').catch(()=>{})">
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
    name: "UserProfile",
    components: { AllPostComments, AddPostComment, AllTags },
    data() {
    return {
      loggedUser: false,
      allUserPosts: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allPostComments: [],
      allPostTags: [],
      allTagsDialog: false,
      postId: 0,
      searchInput: "",
      username: "",
      name: "",
      surname: "",
      biography: "",
      webSite: "",
      space: " "
    };
  },
  created() {
        this.username = this.$route.params.username;
    },

  methods: {
    checkLoggedUser() {
      if (getId().length != 0) {
        this.loggedUser = true;
      }
    },
    loadUserProfileData(){
        axios.get("http://localhost:8081/api/user/regular-user-by-username/"+this.username, {
            headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.userData = response.data;
          this.name = this.userData.name;
          this.surname = this.userData.surname;
          this.biography = this.userData.biography;
          this.webSite = this.userData.webSite;
        }).catch(error => {
                if(error.response.status === 404){
                    console.log = "Account with that username doesn't exist!";
                }
        });
    },
    loadAllUserPosts() {
      axios.get("http://localhost:8081/api/media-content/search-user/"+this.username, {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.allUserPosts = response.data;
        })
        .catch((error) => {
          if (error.response.status === 500) {
            console.log("Internal server error");
          }
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
  },
  mounted() {
    this.checkLoggedUser();
    this.loadAllUserPosts();
    this.loadUserProfileData();
  },
}

</script>

<style>

</style>