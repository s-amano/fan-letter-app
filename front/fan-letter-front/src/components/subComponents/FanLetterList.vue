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
              <v-dialog v-model="dialog" :retain-focus="false">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    class="mr-4"
                    dark
                    color="primary"
                    v-bind="attrs"
                    v-on="on"
                    >削除</v-btn
                  >
                </template>
                <v-card>
                  <v-card-title class="headline grey lighten-2">
                    本当に削除しますか？
                  </v-card-title>
                  <v-card-text>
                    {{
                      fanletter.from
                    }}さんがくれたあなたへのファンレターを本当に削除しますか？
                  </v-card-text>

                  <v-divider></v-divider>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="primary"
                      @click="
                        deleteFanletter(fanletter.id, fanletter.text, index)
                      "
                    >
                      削除
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
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
    dialog: false,
  }),
  mounted() {
    axios
      .get(
        "https://ujcabfcro1.execute-api.ap-northeast-1.amazonaws.com/dev/get"
      )
      .then((response) => (this.fanletters = response.data))
      .catch((error) => console.log(error));
  },
  methods: {
    deleteFanletter(id, text, index) {
      const lambdaApi =
        "https://0lg3gcph42.execute-api.ap-northeast-1.amazonaws.com/dev/delete/" +
        id +
        "/" +
        text;

      axios
        .delete(lambdaApi)
        .then(() => {
          console.log("削除成功");
          this.dialog = false;
          this.fanletters.splice(index, 1);
        })
        .catch(() => {
          console.log("失敗");
        });
    },
  },
};
</script>
