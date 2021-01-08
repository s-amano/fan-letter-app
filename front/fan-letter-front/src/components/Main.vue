<template>
  <div>
    <v-row>
      <v-col>
        <Form v-if="page === 1" @viewList="viewList" />
        <FanLetterList v-if="page === 2" />
      </v-col>
    </v-row>
    <v-row class="buttons">
      <v-col>
        <v-btn class="mr-4" @click="viewForm">ファンレターを書く</v-btn>
        <v-btn class="mr-4" @click="viewList">ファンレターを見る</v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import Form from './subComponents/Form';
import FanLetterList from './subComponents/FanLetterList';

export default {
  components: {
    Form,
    FanLetterList,
  },
  data: () => ({
    page: 1,
  }),
  methods: {
    viewForm() {
      this.page = 1;
    },
    viewList() {
      this.page = 2;
    },
  },
};
</script>

<style>
.buttons {
  text-align: center;
}
</style>
