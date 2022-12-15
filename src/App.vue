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
    <AddProfile @add-profile="addProfile" :v-model="profiles" />
  </div>
  <div v-if="num === 5" class="container2">
    <AddCard />
  </div>
  <div v-if="num === 2" class="container2">
    <CreateTask @create-task="addTask" />
  </div>
</template>

<script>
import Tasks from "./components/Tasks.vue";
import Button from "./components/Button";
import AddProfile from "./components/AddProfile";
import AddCard from "./components/AddCard";
import CreateTask from "./components/CreateTask.vue";

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
    };
  },
  methods: {
    addTask(task) {
      this.tasks = [...this.tasks, task];
    },
    deleteTask(id) {
      this.tasks = this.tasks.filter((task) => task.id !== id);
    },
    AddProfile(profile) {
      this.profiles = [...this.profiles, profile];
      console.log(profile.fname);
    },
  },
  created() {
    this.tasks = [
      {
        id: 1,
        task: "Task1",
        shop: "Supreme",
        product: "Air Force 1",
        size: "43",
        color: "white",
        profile: "profile1",
        card: "card1",
        status: "paused",
        reminder: true,
      },
      {
        id: 2,
        task: "Task2",
        shop: "Nike",
        product: "Jordan 1",
        size: "43",
        color: "blue",
        profile: "profile1",
        card: "card1",
        status: "paused",
        reminder: true,
      },
    ];
    this.profiles = [
      {
        fname: "a",
        lname: "b",
        email: "x",
        phone: "x",
        address: "x",
        address2: "x",
        postcode: "x",
        city: "x",
        county: "x",
        country: "x",
      },
      {
        fname: "b",
        lname: "b",
        email: "y",
        phone: "y",
        address: "y",
        address2: "y",
        postcode: "y",
        city: "y",
        county: "y",
        country: "y",
      },
    ];
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
  height: 600px;
  border: 1px solid black;
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
