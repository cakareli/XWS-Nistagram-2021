<template>
  <v-app class="grey lighten-4" width="800px">

    <v-app-bar app height="45" color="grey lighten-3">
      <v-app-bar-nav-icon @click.stop = "drawer=!drawer" class="ml-5"/>
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
    
        <v-navigation-drawer
          v-model="drawer"
          absolute
          class="grey lighten-4"
          fixed     
        >
          <v-list>
            <v-list-item-group active-class="deep-purple--text text--accent-4">
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
                  Privacy</v-list-item-title
                >
              </v-list-item>

              <v-list-item>
                <v-list-item-title>
                  <v-icon small>mdi-bookmark</v-icon>
                  Saved Posts</v-list-item-title
                >
              </v-list-item>

              <v-list-item>
                <v-list-item-title>
                  <v-icon small>mdi-text-box-multiple</v-icon>
                  Verify</v-list-item-title
                >
              </v-list-item>

              <v-list-item @click="logout">
                <v-list-item-title>
                  <v-icon small>mdi-logout</v-icon>
                  Log out</v-list-item-title
                >
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-navigation-drawer>



    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">         
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="post in allUserPosts" :key="post.Username">
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
                      >Like</v-btn
                    >
                    <v-btn
                      x-small
                      @click="dislikePost(post.Id)"
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

            <v-btn class= "mx-2" @click="$router.push('/search').catch(()=>{})">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2"
              @click="$router.push('/new-post').catch(()=>{})"
            >
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn class= "mx-2">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            <v-btn
              class= "mx-2" @click="$router.go()"
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
import axios from "axios";
import {
  getId,
  getToken,
  getUsername,
  removeToken,
} from "../security/token.js";
import AllPostComments from "../components/AllPostComments.vue";
import AddPostComment from "../components/AddPostComment.vue";
import AllTags from "../components/AllTags.vue";
import AllHashtags from "../components/AllHashtags.vue";


export default {
  name: "Account",
  components: { AllPostComments, AddPostComment, AllTags, AllHashtags },
  data() {
    return {
      drawer: false,
      name: "",
      surname: "",
      username: "",
      biography: "",
      website: "",
      allUserPosts: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allHashtagsDialog: false,
      allPostHashtags: [],
      allPostComments: [],
      allPostTags: [],
      allTagsDialog: false,
      postId: 0
    };
  },
  methods: {
    loadRegisteredUser() {
      axios
        .get("http://localhost:8081/api/user/logged-user/" + getId(), {
          headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.regularUser = response.data;
          this.name = this.regularUser.Name + " ";
          this.surname = this.regularUser.Surname;
          this.username = this.regularUser.Username;
          this.biography = this.regularUser.Biography;
          this.website = this.regularUser.WebSite;
        })
        .catch((error) => {
          if (error.response.status === 500) {
            this.snackbarText = "Internal server error occurred!";
            this.snackbar = true;
          }
          if (error.response.status === 401) {
            this.snackbar = true;
            this.snackbarText =
              "You are unauthorized to get patient informations!";
          }
        });
    },
    logout() {
      removeToken();
      this.$router.push("/").catch(() => {});
    },
    loadUsersPosts() {
      axios
        .get(
          "http://localhost:8081/api/media-content/regular-user-posts/" +
            getUsername(),
          {
            headers: {
              Authorization: "Bearer " + getToken(),
            },
          }
        )
        .then((response) => {
          this.allUserPosts = response.data;
        })
        .catch((error) => {
          if (error.response.status === 500) {
            console.lof("Internal server error");
          }
        });
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
    }
  },
  mounted() {
    this.loadRegisteredUser();
    this.loadUsersPosts();
  },
};
</script>

<style>
</style>