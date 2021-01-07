<template>
  <ValidationObserver ref="obs" v-slot="{ invalid }">
    <form>
      <ValidationProvider v-slot="{ errors, valid }" name="id" rules="required">
        <v-text-field v-model="ID" :error-messages="errors" label="ID" required :success="valid"></v-text-field>
      </ValidationProvider>

      <ValidationProvider v-slot="{ errors, valid }" name="テキスト" rules="required">
        <v-text-field v-model="text" :error-messages="errors" label="テキスト" required :success="valid"></v-text-field>
      </ValidationProvider>

      <v-btn class="mr-4" @click="submit" :disabled="invalid" color="success">送信</v-btn>
    </form>
  </ValidationObserver>
</template>

<script>
import { required } from 'vee-validate/dist/rules';
import { localize, extend, ValidationObserver, ValidationProvider } from 'vee-validate';
import ja from 'vee-validate/dist/locale/ja.json';

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
  }),
  methods: {
    submit() {
      this.$refs.obs.validate().then((result) => {
        console.log('submit', result);
      });
    },
  },
};
</script>
