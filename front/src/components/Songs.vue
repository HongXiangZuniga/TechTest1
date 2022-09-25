<template>
  <div class="row">
    <div class="col"></div>
    <div class="col">
      <label for="band">Nombre de la banda</label><br>
      <input type="text" id="band" name="bandName" v-model="band"><br>
      <br>
      <button type="button" class="btn btn-primary" v-on:click='getTracks(this.band)'
        :disabled="isSpinner">Buscar</button>
      <br>
      <br>
      <div class="spinner-border" role="status" v-if="isSpinner">
        <span class="sr-only">Loading...</span>
      </div>
      <div class="row"  v-if="isActive">
        <div class="col">Albumes Totales:{{ this.totalCollection}}</div>
        <div class="col">Canciones Totales:{{ this.totalSongs}}</div>
      </div>
      <BandTable v-if="isActive" v-bind:songs="this.songs"></BandTable>
    </div>
    <div class="col"></div>
  </div>
</template>

<script>
import axios from "axios";
import BandTable from "./BandTable.vue";
export default {
  name: 'songs-view',
  data() {
    return {
      isSpinner: false,
      isActive: false,
      songs: null,
      totalCollection: 0,
      totalSongs: 0
    }
  },
  components: {
    BandTable,
  },
  props: {
  },
  methods: {
    getTracks: async function (band) {
      this.isActive = false;
      this.isSpinner = true;
      console.log(process.env.VUE_APP_URL_API +"?name="+ encodeURIComponent(band))
      axios
        .get(process.env.VUE_APP_URL_API +"?name="+ encodeURIComponent(band))
        .then(response => {
          this.isActive = true;
          this.songs = response.data.canciones;
          this.isSpinner = false;
          this.totalCollection = response.data.total_albumes
          this.totalSongs = response.data.total_canciones
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
