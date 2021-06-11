<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45" color="#A29D9C" width="600px" >
            <v-app-bar app height="45">
              <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
              <v-row >
                <v-col>
                  <v-toolbar-title>
                    <span outlined class="font-weight-light">NISTA</span>
                    <span>GRAM</span>
                  </v-toolbar-title>
                </v-col>
                <v-spacer></v-spacer>
                <v-col>
                  <v-btn outlined class="mx-5 white" small>
                    <v-icon>mdi-send</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-app-bar>
            <v-navigation-drawer v-model="drawer" absolute  temporary height="500" class="grey lighten-4">
              <v-list nav dense>
                <v-list-item-group
                  v-model="group"
                  active-class="deep-purple--text text--accent-4">
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
                      Privacy</v-list-item-title>
                  </v-list-item>

                  <v-list-item>
                    <v-list-item-title>
                      <v-icon small>mdi-bookmark</v-icon>
                      Saved Posts</v-list-item-title>
                  </v-list-item>

                  <v-list-item>
                    <v-list-item-title>
                      <v-icon small>mdi-text-box-multiple</v-icon>
                      Verify</v-list-item-title>
                  </v-list-item>

                </v-list-item-group>
              </v-list>
            </v-navigation-drawer>
          </v-toolbar>
          <v-toolbar id="guestToolbar" ref="guestToolbar" height="35" class="grey lighten-4" width="600px" hidden>
            <v-app-bar app>
              <v-row align="center" justify="space-around">
                <v-col>
                  <v-btn
                    width="250px"
                    height="35px"
                    @click="$router.push('/registration')"
                    >Register</v-btn
                  >
                </v-col>
                <v-col>
                  <v-btn
                    width="250px"
                    height="35px"
                    @click="$router.push('/login')"
                    >Login</v-btn
                  >
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-row>
              <v-col width="200px">     
              </v-col>
              <v-col>
                <v-btn height="500">

                </v-btn>
              </v-col>
            </v-row>
          </v-container>
          <v-bottom-navigation height="35" width="600px">
            <v-btn value="home">
              <v-icon>mdi-home</v-icon>
            </v-btn>

            <v-btn value="search">
              <v-icon>mdi-magnify</v-icon>
            </v-btn>

            <v-btn value="add">
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>

            <v-btn value="notification">
              <v-icon>mdi-bell-ring</v-icon>
            </v-btn>

            <v-dialog>
              <template v-slot:activator="{ on }">
                <v-btn
                  value="profile"
                  v-on="on"
                  @click="$router.push('/account')"
                >
                  <v-icon>mdi-account</v-icon>
                </v-btn>
              </template>
              <v-card>
                <v-card-title>
                  <h2>Account</h2>
                </v-card-title>
              </v-card>
            </v-dialog>
          </v-bottom-navigation>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
  </v-app>
  <!-- <v-container>
    <v-row>
      <v-col>
        <h1>Account</h1>
        <v-btn @click="$router.push('/update')">Edit profile</v-btn>
      </v-col>
    </v-row>
  </v-container> -->
</template>

<script>
export default {
  name: "Account",
  data() {
    return{
      drawer : false
    }
  }
};
</script>

<style>
</style>