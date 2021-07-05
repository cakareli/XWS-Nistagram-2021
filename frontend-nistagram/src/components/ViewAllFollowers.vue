<template>
    <v-dialog v-model="viewAllFollowersDialog" width = "400" height="500" persistent>
        <v-layout justify-center>
             <v-flex>
                <v-card color="grey lighten-4" class="pa-3">
                    <v-card-title>
                            <h3>All Followers:</h3>
                    </v-card-title>
                        <v-row v-for= "item in allFollowers" :key = "item.userId">
                            <v-col>
                                <a icon :href="getLink(item.username)" class="ml-3">
                                    @{{item.username}} 
                                </a>
                            </v-col>
                        </v-row>
                    <v-card-actions>
                        <v-btn color="grey lighten-3" @click.native="close">Close</v-btn>
                    </v-card-actions>
                </v-card>
             </v-flex>
        </v-layout>
    </v-dialog>
</template>

<script>

export default {
   name: 'ViewAllFollowers',
    components: {
    },
    props: {
        viewAllFollowersDialog: {
        default: false,
        },
        allFollowers : {},
    },
    data() {
        return {
            
        }
    },
    methods : {
        close() {
        this.$emit('update:viewAllFollowersDialog', false)
        },
        getLink(username){
            return '/user-profile/' + username;
        }
    }
}
</script>
