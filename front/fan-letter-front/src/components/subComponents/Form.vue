<template>
  <ValidationObserver ref="obs" v-slot="{ invalid }">
    <form>
      <ValidationProvider v-slot="{ errors, valid }" name="id" rules="required">
        <v-text-field v-model="ID" :error-messages="errors" label="ID:123" required :success="valid"></v-text-field>
      </ValidationProvider>

      <ValidationProvider v-slot="{ errors, valid }" name="手紙内容" rules="required">
        <v-text-field
          v-model="text"
          :error-messages="errors"
          label="手紙内容:いつもありがとう"
          required
          :success="valid"
        ></v-text-field>
      </ValidationProvider>

      <ValidationProvider v-slot="{ errors, valid }" name="From" rules="required">
        <v-text-field
          v-model="from"
          :error-messages="errors"
          label="From:天野"
          required
          :success="valid"
        ></v-text-field>
      </ValidationProvider>

      <v-btn class="mr-4" @click="submit" :disabled="invalid" color="success">送信</v-btn>
    </form>
  </ValidationObserver>
</template>

<script>
import { required } from 'vee-validate/dist/rules';
import { localize, extend, ValidationObserver, ValidationProvider } from 'vee-validate';
import ja from 'vee-validate/dist/locale/ja.json';
import axios from 'axios';

extend('required', required);
localize('ja', ja);

export default {
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  data: () => ({
    ID: '',
    text: '',
    from: '',
  }),
  methods: {
    submit() {
      this.$refs.obs.validate().then(() => {
        const lambdaApi = 'https://4xwqbgj0ad.execute-api.ap-northeast-1.amazonaws.com/dev/receiver';

        const postData = {
          id: this.ID,
          text: this.text,
          from: this.from,
        };

        console.log(postData);

        const headers = {
          headers: {
            'Content-Type': 'application/json',
          },
        };

        axios
          .post(lambdaApi, postData, headers)
          .then(() => {
            console.log('成功');
            this.viewList();
          })
          .catch(() => {
            console.log('失敗');
          });
      });
    },
    viewList: function() {
      this.$emit('viewList');
    },
  },
};
</script>
