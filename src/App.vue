<template>
  <div class="container">
    <header>
      <Button v-on:Click="num = 1" text="Tasks" color="black" />
      <Button v-on:Click="num = 3" text="Profiles" color="black" />
      <Button v-on:Click="num = 5" text="Cards" color="black" />
    </header>
  </div>
  <div v-if="num === 1" class="container2">
    <div class="container3">
      <table>
        <tr>
          <th style="max-width: 150px; min-width: 150px">Task</th>
          <th style="max-width: 140px; min-width: 140px">Shop</th>
          <th style="max-width: 160px; min-width: 160px">Product</th>
          <th style="max-width: 90px; min-width: 90px">Size</th>
          <th style="max-width: 200px; min-width: 200px">Color</th>
          <th style="max-width: 90px; min-width: 90px">Status</th>
          <th style="max-width: 150px; min-width: 150px">Controls</th>
        </tr>
      </table>
    </div>
    <Tasks @delete-task="deleteTask" :tasks="tasks" />
    <Button
      v-on:Click="num = 2"
      class="BtnCreateTask"
      text="Create Task"
      color="black"
    />
  </div>
  <div v-if="num === 3" class="container2">
    <AddProfile
      :profiles="profiles"
      @add-profile="addProfile"
      @edit-profile="editProfile"
      @delete-profile="deleteProfile"
    />
  </div>
  <div v-if="num === 5" class="container2">
    <AddCard
      :cards="cards"
      @add-card="addCard"
      @edit-card="editCard"
      @delete-card="deleteCard"
    />
  </div>
  <div v-if="num === 2" class="container2">
    <CreateTask @create-task="addTask" :profiles="profiles" :cards="cards" />
  </div>
</template>

<script>
import Tasks from "./components/Tasks.vue";
import Button from "./components/Button";
import AddProfile from "./components/AddProfile";
import AddCard from "./components/AddCard";
import CreateTask from "./components/CreateTask.vue";
import axios from "axios";
export default {
  name: "App",
  components: {
    Tasks,
    Button,
    AddProfile,
    AddCard,
    CreateTask,
  },
  data() {
    return {
      num: 1,
      tasks: [],
      profiles: [],
      cards: [],
    };
  },
  methods: {
    addTask(task) {
      this.tasks = [...this.tasks, task];
    },
    deleteTask(id) {
      this.tasks = this.tasks.filter((task) => task.id !== id);
    },
    addProfile(profile) {
      this.profiles.push(profile);

      //axios.post('http:3000/profile/add', profile)
      // axios.get('http:3000/profile/list')
    },
    editProfile(profile) {
      const index = this.profiles.findIndex(
        (element) => element.id === profile.id
      );
      this.profiles.splice(index, 1, profile);
    },
    deleteProfile(profile) {
      const index = this.profiles.findIndex(
        (element) => element.id === profile.id
      );
      this.profiles.splice(index, 1);
    },
    addCard(card) {
      this.cards.push(card);
    },
    editCard(card) {
      const index = this.cards.findIndex((element) => element.id === card.id);
      this.cards.splice(index, 1, card);
    },
    deleteCard(card) {
      const index = this.cards.findIndex((element) => element.id === card.id);
      this.cards.splice(index, 1);
    },
    async initCards() {
      let response = await axios.get("http://localhost:9000/cards");
      this.cards = [...response?.data];
    },
    async initProfiles() {
      let response = await axios.get("http://localhost:9000/profiles");
      this.profiles = [...response?.data];
    },
    async initTasks() {
      let response = await axios.get("http://localhost:9000/tasks");
      this.tasks = [...response?.data];
    },
  },
  async created() {
    this.initCards();
    this.initProfiles();
    this.initTasks();
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400&display=swap");
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
body {
  font-family: "Poppins", sans-serif;
  background: rgb(31, 31, 31);
}
.container {
  max-width: 800px;
  margin: 90px auto;
  margin-top: 30px;
  margin-bottom: 40px;
  overflow: hidden;
  height: 100px;
  border: 1px solid black;
  padding: 25px;
  border-radius: 5px;
}
.container2 {
  max-width: 1100px;
  margin: 60px auto;
  margin-top: 30px;
  overflow: hidden;
  height: 800px;
  border: 0px solid black;
  padding: 10px;
  border-radius: 5px;
}
.container3 {
  background: cadetblue;
  width: 1020px;
  height: 30px;
  margin: 10px;
  margin-left: 25px;
  margin-right: 25px;
  padding: 5px 10px;
  cursor: pointer;
  border-radius: 10px;
}

.BtnCreateTask {
  position: fixed;
  left: 870px;
  top: 800px;
}

.container3 th {
  text-align: center;
  max-height: 25px;
}

.container3 table {
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
}

.btn {
  display: inline-block;
  background: #000;
  color: #fff;
  border: none;
  padding: 10px 20px;
  margin: 5px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
}
.btn:focus {
  outline: none;
}
.btn:active {
  transform: scale(0.98);
}
.btn-block {
  display: block;
  width: 100%;
}
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
