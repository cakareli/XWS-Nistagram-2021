<template>
  <v-app class="grey lighten-4">
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar height="45" color="#A29D9C" width="600px">
            <v-app-bar app>
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
          </v-toolbar>
          <v-toolbar height="35" class="grey lighten-4" width="600px">
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
              <v-col> </v-col>
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
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {};
  },
};
