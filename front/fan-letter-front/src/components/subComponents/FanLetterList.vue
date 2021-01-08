<template>
  <v-card class="mx-auto" max-width="300" tile>
    <v-list dense>
      <v-subheader>
        <h2>ファンレターリスト</h2>
      </v-subheader>
      <!-- <v-list-item-group color="primary"> -->
      <v-list-item v-for="(fanletter, i) in fanletters" :key="i">
        <v-list-item-content>
          <v-list-item-title v-text="fanletter.text"></v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <!-- </v-list-item-group> -->
    </v-list>
  </v-card>
</template>

<script>
import axios from 'axios';
export default {
  data: () => ({
    fanletters: [],
  }),
  mounted() {
    axios
      .get('https://ujcabfcro1.execute-api.ap-northeast-1.amazonaws.com/dev/get')
      .then((response) => (this.fanletters = response.data))
      .catch((error) => console.log(error));
  },
};
</script>
