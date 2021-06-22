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
    <v-container>
      <v-row justify="center">
        <v-card width="800px" class="pa-12">
          <v-row justify="center">
            <v-list>
              <v-list-item v-for="user in closeFriends" :key="user.username">
                <v-card height="60" width="600" class="ma-3 grey lighten-5">
                  <v-card-title height="60" class="grey lighten-3">
                    <v-row>
                      <v-col>
                        <v-row>
                          <h4>@{{ user.username }}</h4>
                        </v-row>
                        </v-col>
                      <v-spacer />
                      <v-col>
                        <v-btn small @click="loadUserProfile(user.username)">
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

          <v-btn class="mx-2">
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
      closeFriends: [],
    };
  },

  methods: {
    loadAllCloseFriends() {
      axios
        .get("http://localhost:8081/api/follow/close-followers/" + getId())
        .then((response) => {
          console.log(response);
          this.closeFriends = response.data;
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
  },
  mounted() {
    this.loadAllCloseFriends();
  },
};
</script>

<style>
</style>