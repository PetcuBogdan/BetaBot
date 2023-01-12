<template>
  <div :class="[task.reminder ? 'reminder' : '', 'task']">
    <tabel>
      <tr>
        <th class="truncate">{{ task.name }}</th>
        <th>{{ task.shop }}</th>
        <th class="truncate">{{ task.productName }}</th>
        <th>{{ task.size }}</th>
        <th class="truncate">{{ task.color }}</th>
        <th>{{ task.status }}</th>
        <i
          v-if="st === 1"
          @click="stop(task.id)"
          class="fa fa-pause spacePause"
        ></i>
        <i
          v-if="st === 0"
          @click="start(task.id)"
          class="fa fa-play spaceStart"
        ></i>
        <i @click="onDelete(task.id)" class="fas fa-times"></i>
      </tr>
    </tabel>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Task-el",
  props: {
    task: Object,
  },
  data() {
    return {
      st: 0,
    };
  },
  methods: {
    onDelete(id) {
      this.$emit("delete-task", id);
      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/taskDelete", {
          id: id,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    start(id) {
      this.st = 1;
      console.log(id);

      axios
        .post("http://localhost:9000/taskStart", {
          id: id,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    stop(id) {
      this.st = 0;
      console.log(id);

      axios
        .post("http://localhost:9000/taskStop", {
          id: id,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
};
</script>

<style scope>
.spacePause {
  margin-right: 10px;
  color: black;
}
.spaceStart {
  margin-right: 10px;
  color: lightgreen;
}
.fas {
  color: red;
}
.task {
  background: rgb(55, 68, 92);
  height: 45px;
  width: 930px;
  margin-left: 70px;
  margin-right: 70px;
  margin-top: 5px;
  margin-bottom: 5px;
  padding: 10px 20px;
  cursor: pointer;
  border-radius: 10px;
}
.task.reminder {
  border-left: 5px solid green;
}

.task table {
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
}
.task th {
  text-align: left;
  min-width: 140px;
  max-width: 140px;
  height: 25px;
}
.truncate {
  width: 100px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
