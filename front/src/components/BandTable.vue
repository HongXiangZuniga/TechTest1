<template>
  <div class="row">
    <div class="col"></div>
    <div class="col">
      <table class="table table-striped" id="table">
        <thead class="thead-dark">
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Album</th>
            <th scope="col">Tema</th>
            <th scope="col">Preview</th>
            <th scope="col">Lanzamiento</th>
            <th scope="col">Valor(USD)</th>
          </tr>
          <tr v-for="(song, index) in this.songs" :key="index">
            <td>
              {{ song.cancion_id}}
            </td>
            <td>
              {{ song.nombre_album}}
            </td>
            <td>
              {{ song.nombre_tema}}
            </td>
            <td>
              {{song.preview_url}}
            </td>
            <td>
              {{song.fecha_lanzamiento}}
            </td>
            <td>
              {{song.precio.valor}}
            </td>
          </tr>
        </thead>
        <tbody>
        </tbody>
      </table>
    </div>
    <div class="col"></div>
  </div>
</template>
  
<script>
export default {
  name: 'band-table',
  props: {
    songs: Array
  },
  methods: {
  },
  mounted() {
    this.tableData = this.songs
  },
  data() {
    return {
      tableData: this.songs
    }
  }
}

</script>
  
<style>
.table {
  text-align: center;
  border: 3px solid;
  padding-left: 30%;
  padding-right: 30%;
  position: relative;
  width: 70%;
}
</style>
  