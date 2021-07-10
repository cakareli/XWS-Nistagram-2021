<template>
  <v-app class="grey lighten-2">
    <v-app-bar app height="45" class="grey lighten-3 justify-center">
      <v-app-bar-nav-icon @click="$router.push('/account')">
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
          <v-btn outlined class="mx-5 white">
            <v-icon>mdi-send</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-app-bar>
    <v-container v-show="requestExists">
      <v-row justify="center">
        <v-card width="900px" class="pa-12">
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="user in followRequest" :key="user.username">
                <v-card height="60" width="800" class="ma-3 grey lighten-5">
                  <v-card-title height="60" class="grey lighten-3">
                    <v-row>
                      <v-col>
                        <v-row class="ma-1">
                          <h4>@{{ user.username }}</h4>
                        </v-row>
                        </v-col>
                      <v-spacer />
                      <v-col>
                        <v-btn small @click="loadUserProfile(user.username)">
                          <span>View profile</span>
                        </v-btn>
                        <v-btn class="success" small @click="accept(user.userId)">
                          <span>Accept</span>
                        </v-btn>
                        <v-btn class="error ml-2" small @click="reject(user.userId)">
                          <span>Reject</span>
                        </v-btn>
                      </v-col>
                    </v-row>
                  </v-card-title>
                  <v-divider></v-divider>
                </v-card>
              </v-list-item>
            </v-list>
          </v-row>
        </v-card>
      </v-row>
    </v-container>
    <br/><br/>
    <v-container v-show="notificationExists">
      <v-row justify="center">
        <v-card width="900px" class="pa-12">
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="notification in notifications" :key="notification.mediaContentId">
                <v-card height="60" width="800" class="ma-3 grey lighten-5">
                  <v-card-title height="60" class="grey lighten-3">
                    <v-row>
                      <v-col>
                        <v-row class="ma-1">
                          <span>{{ notification.text }}</span>
                        </v-row>
                      </v-col>
                      <v-col v-show="notification.notificationType !== 1">
                        <v-btn small @click="loadContent(notification.mediaContentId, notification.notificationType)">
                          <span>View post</span>
                        </v-btn>
                      </v-col>
                      <v-col v-show="notification.notificationType === 1">
                        <v-btn small @click="loadContent(notification.mediaContentId, notification.notificationType)">
                          <span>View story</span>
                        </v-btn>
                      </v-col>
                    </v-row>
                  </v-card-title>
                  <v-divider></v-divider>
                </v-card>
              </v-list-item>
            </v-list>
          </v-row>
        </v-card>
      </v-row>
    </v-container>

    <v-footer app height="45px" class="grey lighten-3 justify-center">
      <v-container>
        <v-row justify="center">
          <v-btn class="mx-2" @click="$router.push('/').catch(() => {})">
            <v-icon>mdi-home</v-icon>
          </v-btn>

          <v-btn class="mx-2" @click="$router.push('/search').catch(() => {})">
            <v-icon>mdi-magnify</v-icon>
          </v-btn>

          <v-btn
            class="mx-2"
            @click="$router.push('/new-post').catch(() => {})"
          >
            <v-icon>mdi-plus-box</v-icon>
          </v-btn>

          <v-btn class="mx-2" @click="$router.go()">
            <v-icon>mdi-bell-ring</v-icon>
          </v-btn>

          <v-btn class="mx-2" @click="$router.push('/account').catch(() => {})">
            <v-icon>mdi-account</v-icon>
          </v-btn>
        </v-row>
      </v-container>
    </v-footer>
  </v-app>
</template>

<script>
import { getId } from "../security/token.js";
import axios from "axios";
export default {
  data() {
    return {
      followRequest: [],
      notifications: [],
      notificationExists: true,
      requestExists: true
    };
  },

  methods: {
    loadAllFollowRequests() {
      axios
        .get("http://localhost:8081/api/follow/follow-requests/" + getId())
        .then((response) => {
          console.log(response);
          this.followRequest = response.data;
          if(this.followRequest.length === 0){
            this.requestExists = false
          }
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },
    accept(id){
      axios
      .put('http://localhost:8081/api/follow/accept-follow/' + getId() + "/" + id)
      .then((response) => {
          console.log(response);
          this.$router.go();
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },
    reject(id){
      axios
      .put('http://localhost:8081/api/follow/remove-follower/' + getId() + "/" + id)
      .then((response) => {
          console.log(response);
          this.$router.go();
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "Bad request, try again!";
            this.snackbar = true;
          }
        });
    },
    loadUserProfile(username){
      this.$router.push('/user-profile/' + username).catch(()=>{})
    },
    loadNotifications(){
      axios.get('http://localhost:8081/api/media-content/notifications/'+getId())
      .then((response)=>{
        this.notifications = response.data
        console.log(response.data.status)
        if(this.notifications.length === 0){
            this.notificationExists = false
          }
      })
      .catch((error)=>{
        if(error.status === 400){
          console.log("Bad request")
        }
      })
    },
    loadContent(id, contentType){
      this.$router.push('/notification-post/' + id+"/"+contentType).catch(()=>{})
    }
  },
  mounted() {
    this.loadAllFollowRequests();
    this.loadNotifications();
  },
};
</script>

<style>
</style>