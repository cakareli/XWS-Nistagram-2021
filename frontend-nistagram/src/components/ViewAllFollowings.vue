<template>
    <v-dialog v-model="viewAllFollowingsDialog" width = "400" height="300" persistent>
        <v-layout justify-center>
             <v-flex>
                <v-card color="grey lighten-4" class="pa-3">
                    <v-card-title>
                            <h3>All Followings:</h3>
                    </v-card-title>
                        <v-row v-for= "item in allFollowings" :key = "item.userId" @click="$router.push('/user-profile/' + item.userId).catch(()=>{})">
                            <v-col>
                                <a icon :href="getLink(item.userId)" class="ml-3">
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
        viewAllFollowingsDialog: {
            default: false,
        },
        allFollowings : {},
    },
    data() {
        return {
            
        }
    },
    methods : {
        close() {
        this.$emit('update:viewAllFollowingsDialog', false)
        },
        getLink(userId){
            return '/user-profile/' + userId;
        }
    }
}
</script>
