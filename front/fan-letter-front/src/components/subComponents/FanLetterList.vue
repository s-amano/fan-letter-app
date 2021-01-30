<template>
  <v-card class="mx-auto" max-width="700" tile>
    <v-list two-line>
      <v-subheader>
        <h2>ファンレターリスト</h2>
      </v-subheader>
      <v-list-item-group color="primary">
        <template v-for="(fanletter, index) in fanletters">
          <v-list-item :key="`first-${index}`">
            <v-list-item-content>
              <v-list-item-title v-text="fanletter.text"></v-list-item-title>
              <v-list-item-subtitle v-if="fanletter.from == ''" v-text="'匿名'"
                >より</v-list-item-subtitle
              >
              <v-list-item-subtitle v-else v-text="fanletter.from + 'より'"
                >より</v-list-item-subtitle
              >
            </v-list-item-content>
            <v-list-item-action>
              <v-list-item-action-text
                v-text="fanletter.post_at"
              ></v-list-item-action-text>
            </v-list-item-action>
          </v-list-item>
          <v-divider
            v-if="index < fanletters.length - 1"
            :key="`second-${index}`"
          ></v-divider>
        </template>
      </v-list-item-group>
    </v-list>
  </v-card>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    fanletters: [],
  }),
  mounted() {
    axios
      .get(
        "https://ujcabfcro1.execute-api.ap-northeast-1.amazonaws.com/dev/get"
      )
      .then((response) => (this.fanletters = response.data))
      .catch((error) => console.log(error));
  },
};
</script>
