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
                <v-col>
                  <br>
                  <br>
                  <v-btn @click="viewFollowers">Followers: {{this.followersNumber}}</v-btn>
                </v-col>
                <v-col>
                  <br>
                  <br>
                  <v-btn @click="viewFollowings">Following: {{this.followingsNumber}}</v-btn>
                </v-col>
            </v-row>
            <v-row justify="center" v-show="!loggedUserProfile">
                <v-btn  class="mx-5" @click="followUser" v-show="!followed">Follow</v-btn>
                <v-btn  class="mx-5" @click="unfollowUser" v-show="followed">Unfollow</v-btn>
                <v-btn  class="mx-5" @click="addToClose" v-show="!isClose">Add To Close Friends</v-btn>
                <v-btn  class="mx-5" @click="removeFromClose" v-show="isClose">Remove From Close Friends</v-btn>
                <v-btn  class="mx-5" @click="muteUser" v-show="!muted">Mute</v-btn>
                <v-btn  class="mx-5" @click="unmuteUser" v-show="muted">Unmute</v-btn>
                <v-btn  class="mx-5" @click="blockUser" v-show="!blocked">Block</v-btn>
                <v-btn  class="mx-5" @click="unblockUser" v-show="blocked">Unblock</v-btn>
            </v-row>
            <br/><br/>
            <v-row justify="center" v-show="!loggedUserProfile">
              <v-btn  class="mx-5" @click="addToClose" v-show="!isClose">Add To Close Friends</v-btn>
              <v-btn  class="mx-5" @click="removeFromClose" v-show="isClose">Remove From Close Friends</v-btn>
              <v-btn  class="mx-5" @click="turnNotificationsOn" v-show="!notification ">Turn notifications on</v-btn>
              <v-btn  class="mx-5" @click="turnNotificationsOff" v-show="notification ">Turn notifications off</v-btn>
            </v-row>
        </v-card>
      </v-row> 
    </v-container>
    <br><br>
    <v-container v-show="adminLogged">
      <br><br>
      <v-row justify="center">
        <v-card width="800px" class=" grey lighten-4">
            <v-row justify="center">
              <v-btn width="800" class="error">Remove Profile</v-btn>
            </v-row>
        </v-card>
      </v-row>
      <br><br>
    </v-container>


    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">
            <v-row justify="center" >
            <v-list>
              <v-list-item v-for="post in allUserPosts" :key="post.Username">
                <v-card height="800" width="550" class="ma-3 grey lighten-5">
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
                  <v-row  class="ma-2" height="5">
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

            <v-btn class= "mx-2" @click="$router.push('/notifications').catch(()=>{})" v-show="loggedUser">
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
    <ViewAllFollowers :viewAllFollowersDialog.sync="viewAllFollowersDialog" :allFollowers="allFollowers"/>
    <ViewAllFollowings :viewAllFollowingsDialog.sync="viewAllFollowingsDialog" :allFollowings="allFollowings"/>
    <ReportPost :reportPostDialog.sync="reportPostDialog" :reportedPost="reportedPost"/>

    </v-app>
</template>

<script>

import { getId, getUsername, getToken } from "../security/token.js";
import AllPostComments from "../components/AllPostComments.vue";
import AddPostComment from "../components/AddPostComment.vue";
import AllTags from "../components/AllTags.vue";
import axios from "axios";
import AllHashtags from "../components/AllHashtags.vue";
import ViewAllFollowers from "../components/ViewAllFollowers.vue";
import ViewAllFollowings from "../components/ViewAllFollowings.vue";
import ReportPost from "../components/ReportPost.vue"

export default {
    name: "UserProfile",
    components: { AllPostComments, AddPostComment, AllTags, AllHashtags, ViewAllFollowers, ViewAllFollowings, ReportPost },
    data() {
    return {
      loggedUser: false,
      loggedUserProfile: false,
      allUserPosts: [],
      allPostCommentsDialog: false,
      addPostCommentDialog: false,
      allHashtagsDialog: false,
      allPostComments: [],
      allPostTags: [],
      allPostHashtags: [],
      allTagsDialog: false,
      postId: 0,
      searchInput: "",
      id: "",
      username: "",
      name: "",
      surname: "",
      biography: "",
      webSite: "",
      privacyType: 0,
      space: " ",
      isPrivate: false,
      followersNumber: 0,
      followingsNumber: 0,
      blocked: false,
      blockedUsers: [],
      mutedUsers: [],
      closeUsers: [],
      followed: false,
      followedUsers: [],
      allFollowers: "",
      allFollowings: "",
      viewAllFollowingsDialog: false,
      viewAllFollowersDialog: false,
      isClose: false,
      muted: false,
      reportPostDialog: false,
      reportedPost: {},
      adminLogged: false,
      notification: false,
      notificationUsers: [],
    };
  },
  created() {
        this.username = this.$route.params.username;
    },

  methods: {
    reportPost(post){
      this.reportPostDialog = true;
      this.reportedPost = post;
    },
    viewFollowers(){
      this.viewAllFollowersDialog = true;
    },
    viewFollowings(){
      this.viewAllFollowingsDialog = true;
    },
    checkLoggedUser() {
      if (getToken() != null) {
        this.loggedUser = true;
      }
    },
    followUser(){
        if(this.privacyType === 0){
          this.isPrivate = false
        }else{
          this.isPrivate = true
        }
        let newFollowDTO = {
          followerId : getId(),
          followedId : this.id,
          isPrivate : this.isPrivate
        }
        axios.post("http://localhost:8081/api/follow/follow",
                  newFollowDTO)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    unfollowUser(){
        axios.post("http://localhost:8081/api/follow/remove-following/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    muteUser(){
        axios.put("http://localhost:8081/api/follow/mute-following/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    unmuteUser(){
        axios.put("http://localhost:8081/api/follow/unmute-following/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    addToClose(){
        axios.put("http://localhost:8081/api/follow/add-to-close/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    removeFromClose(){
        axios.put("http://localhost:8081/api/follow/remove-from-close/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },

    blockUser(){
      axios.post("http://localhost:8081/api/follow/block-user/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    unblockUser(){
      axios.post("http://localhost:8081/api/follow/unblock-user/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    loadUserProfileData(){
        axios.get("http://localhost:8081/api/user/regular-user-by-username/" + this.username, {
            headers: {
            Authorization: "Bearer " + getToken(),
          },
        })
        .then((response) => {
          this.userData = response.data;
          this.id = this.userData._id
          this.name = this.userData.name;
          this.surname = this.userData.surname;
          this.biography = this.userData.biography;
          this.webSite = this.userData.webSite;
          this.privacyType = this.userData.profilePrivacy.PrivacyType
          if(!this.loggedUser){
             this.loggedUserProfile = true;
          }else{
            if(this.id === getId()){
            this.loggedUserProfile = true;
            }
          }
          

          axios.get("http://localhost:8081/api/follow/followers/"+this.id, {
            headers: {
            Authorization: "Bearer " + getToken(),
          }
          })
          .then(response => {
            if(response.data === null){
              this.followersNumber = 0
            
            }else{
              this.followersNumber = response.data.length
              this.allFollowers = response.data;
            }
            console.log(response);
            this.followedUsers = response.data
            for(let i = 0; i<this.followedUsers.length; i++){
                if(getUsername() === this.followedUsers[i].username){
                  this.followed = true;
                }
            }
          })

          axios.get("http://localhost:8081/api/follow/followings/"+this.id, {
            headers: {
            Authorization: "Bearer " + getToken(),
          }
          })
          .then(response => {
            if(response.data === null){
              this.followingsNumber = 0
              
            }else{
              this.followingsNumber = response.data.length
              this.allFollowings = response.data
            }
          })

          this.loadAllUserPosts()
          this.checkNotificationFollowers()
      })
        
    },
    loadAllUserPosts() {
      let userFound = false;
      if(this.privacyType === 1){
        if(this.allFollowers!==[]){
          for(let i = 0; i<this.allFollowers.length; i++){
            if(getUsername().equals(this.allFollowers[i].username)){
              userFound = true;

            } 
          }
        }
      } else if (this.privacyType === 0 || userFound){
        axios.get("http://localhost:8081/api/media-content/regular-user-posts/"+ this.username, {
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
      }
      
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
    viewAllHashtags(allPostHashtags) {
      this.allPostHashtags = allPostHashtags;
      this.allHashtagsDialog = true;
    },
    checkBlockedUser(){
      axios
        .get("http://localhost:8081/api/follow/blocked-users/" + getId())
        .then((response) => {
          console.log(response);
          
          this.blockedUsers = response.data;
          for(let i = 0; i<this.blockedUsers.length; i++){
              if(this.username === this.blockedUsers[i].username){
                this.blocked = true;
              }
          }
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },

    checkMutedUser(){
      axios
        .get("http://localhost:8081/api/follow/muted-users/" + getId())
        .then((response) => {
          console.log(response);
          
          this.mutedUsers = response.data;
          for(let i = 0; i<this.mutedUsers.length; i++){
              if(this.username === this.mutedUsers[i].username){
                this.muted = true;
              }
          }
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },

    checkCloseFriends(){
      axios
        .get("http://localhost:8081/api/follow/close-followers/" + getId())
        .then((response) => {
          console.log(response);
          
          this.closeUsers = response.data;
          for(let i = 0; i<this.closeUsers.length; i++){
              if(this.username === this.closeUsers[i].username){
                this.isClose = true;
              }
          }
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },
    turnNotificationsOn(){
      axios.put("http://localhost:8081/api/follow/notifications-on/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    turnNotificationsOff(){
      axios.put("http://localhost:8081/api/follow/notifications-off/" + getId() + "/" + this.id)
        .then(response => {
            console.log(response)
            this.$router.go()
        }).catch(error => {
          if(error.response.status === 400){
            console.log("Bad request")
          }
        })
    },
    checkNotificationFollowers(){
      axios.get("http://localhost:8081/api/follow/followers-with-notifications/" + this.id)
        .then((response) => {
          console.log(response);
          
          this.notificationUsers = response.data;
          for(let i = 0; i<this.notificationUsers.length; i++){
            alert(getId() +"==="+ this.notificationUsers[i])
              if(getId() === this.notificationUsers[i]){
                alert("Prolazi")
                this.notification = true;
              }
          }
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },
  },
  mounted() {
    this.checkLoggedUser();
    this.loadUserProfileData();
    this.checkBlockedUser();
    this.checkMutedUser();
    this.checkCloseFriends();
  },
}

</script>

<style>

</style>