<template>
    <v-dialog v-model="allPostCommentsDialog" width = "500" persistent>
        <v-layout justify-center>
             <v-flex>
            <v-card class="pa-3">
                <v-card-title>
                        <h4>All comments:</h4>
                </v-card-title>
                    <v-col v-for= "comment in allComments" :key = "comment.id">
                        <span>@{{comment.username}}</span>
                        <v-textarea outlined v-model="comment.text" prepend-icon="mdi-comment" rows="2" no-resize readonly></v-textarea>
                    </v-col>
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
    name: 'AllPostComments',
    components: {

    },
    props: {
        allPostCommentsDialog: {
        default: false,
      },
    },
    data() {
        return {
            allComments: [
                {id: "1", username: "denisfruza", text: "Lep si batoo <3"},
                {id: "2", username: "dbulaja98", text: "mmmmmmmmm"},
                {id: "3", username: "jarulja", text: "nice"},
                {id: "4", username: "cakinjoo", text: "hvala svima ljubim vas!!"},
            ]
        }
    },
    mounted() {

    },
    methods: {
        close() {
            this.$emit('update:allPostCommentsDialog', false)
        },
    }
}
</script>

