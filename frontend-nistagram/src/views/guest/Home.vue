<template>
    <v-app class="grey lighten-4">
        <v-div id="app">
            <v-div class="app-phone">
                <v-div class="logo">
                    <LogoToolbar></LogoToolbar>
                </v-div>
                <v-div class="feed-backgorund"></v-div>
                <v-div class="lower-toolbar">
                    
                </v-div>
            </v-div>
        </v-div>
    </v-app>
    <!-- <v-div id="app">
        <v-div class="app-phone">
            <v-div class="logo">
                <v-img src="../../assets/logoedited2.png"></v-img>
                <button class="send-button"> 
                    <img src="../../assets/send.png" height="50" width="50">
                </button> 
            </v-div>
            <v-div class="row-home">
                <v-button class="column-home" @click="$router.push('/registration')">Register</v-button>
                <v-button class="column-home">Login</v-button>
            </v-div>
            <v-div class="feed-backgorund"></v-div>
            <v-div class="lower-toolbar">
                <v-button>
                    <v-img src="../../assets/home2.png" width="35" height="35"></v-img>
                </v-button>
                <v-button>
                    <v-img src="../../assets/search.png" width="35" height="35"></v-img>
                </v-button>
                <v-button>
                    <v-img src="../../assets/add.png" width="35" height="35"></v-img>
                </v-button>
                <v-button>
                    <v-img src="../../assets/profile2.png" width="35" height="35"></v-img>
                </v-button>
                <v-button>
                    <v-img src="../../assets/notification2.png" width="35" height="35"></v-img>
                </v-button>
            </v-div>
        </v-div>
    </v-div> -->
</template>

<script>
import LogoToolbar from '../../components/LogoToolbar';

export default {
  name: 'Home',
  components: {
      LogoToolbar
  }
//   methods:{

//     LogoToolbar      register(){
//            return null;
//       }
//   }
    
}
</script>

<style lang="scss" src="../../style/home.scss">

</style>