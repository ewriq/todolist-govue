
<template>
  <div>
    <div class="box">
      <form @submit.prevent="add">
        <div class="field has-addons">
          <div class="control">
            <input class="input" type="text" v-model="vaule" placeholder="Listeye ekle">
          </div>
          <button class="button is-success" type="submit">Gönder</button>
        </div>
      </form>
    </div>
    <div class="box">
      <table class="table">
      <thead>
        <tr>
          <th>Güncelle</th>
          <th>Sil</th>
          <th>Durum</th>
          <th>Todo</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="data in datas" :key="data.no">
          <td><button class="button is-link is-small" @click="upt(data.token)">🎀</button></td>
          <td><button class="button is-danger is-small" @click="del(data.token)">❌</button></td>
          <td>{{ data.can }}</td>
          <td>{{ data.title }}</td>
        </tr>
      </tbody>
    </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      datas: [],
      vaule: ''
    };
  },
  created() {
    this.fetchdatas();
  },

  methods: {
    fetchdatas() {
      axios
        .get("http://127.0.0.1:3000/todos")
        .then((response) => {
          this.datas = response.data.data;
        })
        .catch((error) => {
          console.error(error);
        });
    },

    add() {
      axios
        .post("http://127.0.0.1:3000/add", {
          title: this.vaule,
        })
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          console.error("saaaaaaas", error);
        });
      window.location.href = "/"
    },
  },
};
</script>

<script setup>
function del(message) {
  axios.get(`http://127.0.0.1:3000/del/${message}`)
    .then(response => {
      this.datas = response.data.data;
    })
    .catch(error => {
      console.error(error);
    });
  alert("Başarılı sıkıs")
  window.location.href = "/"
}
function upt(message) {
  axios.get(`http://127.0.0.1:3000/upt/${message}`)
    .then(response => {
      this.datas = response.data.data;
    })
    .catch(error => {
      console.error(error);
    });
  alert("Başarılı sıkıs")
  window.location.href = "/"
}

</script>
