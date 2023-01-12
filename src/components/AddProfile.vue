<template>
  <div>
    <select
      class="btnProfile btn-blockProfile ;"
      id="profileListId"
      @change="test()"
    >
      <option value="0" selected>Modify profile ...</option>
      <option
        v-for="profile in profiles"
        :key="profile.id"
        :value="JSON.stringify(profile)"
      >
        {{ profile.fname }}
      </option>
    </select>
  </div>
  <form @submit.prevent="onSubmit" class="add-form">
    <div class="form-control">
      <label>First name</label>
      <input
        required
        type="text"
        v-model="profile.fname"
        name="Fname"
        placeholder="First name"
      />
    </div>
    <div class="form-control">
      <label>Last name</label>
      <input
        required
        type="text"
        v-model="profile.lname"
        name="Lname"
        placeholder="Last name"
      />
    </div>
    <div class="form-control">
      <label>Email</label>
      <input
        required
        type="text"
        v-model="profile.email"
        name="Email"
        placeholder="Email"
      />
    </div>
    <div class="form-control">
      <label>Phone </label>
      <input
        required
        type="text"
        v-model="profile.phone"
        name="Phone"
        placeholder="Phone number"
      />
    </div>
    <div class="form-control">
      <label>Address</label>
      <input
        required
        type="text"
        v-model="profile.address"
        name="Adress"
        placeholder="Adress"
      />
    </div>
    <div class="form-control">
      <label>Address2</label>
      <input
        required
        type="text"
        v-model="profile.address2"
        name="Address2"
        placeholder="Adress2"
      />
    </div>
    <div class="form-control">
      <label>Postcode</label>
      <input
        required
        type="text"
        v-model="profile.postcode"
        name="Postcode"
        placeholder="Postcode"
      />
    </div>
    <div class="form-control">
      <label>City</label>
      <input
        required
        type="text"
        v-model="profile.city"
        name="City"
        placeholder="City"
      />
    </div>
    <div class="form-control">
      <label>County</label>
      <input
        type="text"
        v-model="profile.county"
        name="County"
        placeholder="County"
      />
    </div>
    <div class="form-control">
      <label>Country</label>
      <CountrySelect v-model="profile.country" name="country" class="drop" />
    </div>
    <input
      v-if="!selected"
      type="submit"
      value="Save Profile"
      class="btnProfile btn-blockProfile ;"
    />
  </form>
  <div class="inner">
    <button
      class="btnProfile btn-blockMod ;"
      v-if="selected"
      @click.stop="onEdit"
    >
      Save Edit Profile
    </button>
    <button
      class="btn-blockMod btnProfileDelete btn-del"
      v-if="selected"
      @click.stop="profileDelete"
    >
      Delete Profile
    </button>
  </div>
</template>

<script>
import CountrySelect from "./CountrySelect.vue";
import axios from "axios";

var IdProfileCount = 1;

export default {
  name: "addProfile",
  props: {
    profiles: Array,
  },
  components: {
    CountrySelect,
  },

  data() {
    return {
      profile: {
        id: "",
        lname: "",
        fname: "",
        email: "",
        phone: "",
        address: "",
        address2: "",
        postcode: "",
        city: "",
        county: "",
        country: "",
      },
      selected: 0,
    };
  },
  methods: {
    test() {
      let x = document.getElementById("profileListId").value;
      let testObject = JSON.parse(x);
      this.profile = { ...testObject };
      this.selected = this.profile.id;
    },
    onSubmit() {
      this.$emit("add-profile", {
        ...this.profile,
        id: IdProfileCount,
      });

      console.log(this.profile.country);
      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/profileAdd", {
          id: IdProfileCount,
          lname: this.profile.lname,
          fname: this.profile.fname,
          email: this.profile.email,
          phone: this.profile.phone,
          address: this.profile.address,
          address2: this.profile.address2,
          postcode: this.profile.postcode,
          city: this.profile.city,
          county: this.profile.county,
          country: this.profile.country,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });

      IdProfileCount += 1;

      this.profile = {
        lname: "",
        fname: "",
        email: "",
        phone: "",
        address: "",
        address2: "",
        postcode: "",
        city: "",
        county: "",
        country: "",
      };
    },
    onEdit() {
      this.$emit("edit-profile", this.profile);
      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/profileEdit", {
          id: this.profile.id,
          lname: this.profile.lname,
          fname: this.profile.fname,
          email: this.profile.email,
          phone: this.profile.phone,
          address: this.profile.address,
          address2: this.profile.address2,
          postcode: this.profile.postcode,
          city: this.profile.city,
          county: this.profile.county,
          country: this.profile.country,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
      this.profile = {
        lname: "",
        fname: "",
        email: "",
        phone: "",
        address: "",
        address2: "",
        postcode: "",
        city: "",
        county: "",
        country: "",
      };
      this.selected = 0;
      document.getElementById("profileListId").value = 0;
    },
    profileDelete() {
      this.$emit("delete-profile", this.profile);
      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/profileDelete", {
          id: this.profile.id,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
      this.profile = {
        lname: "",
        fname: "",
        email: "",
        phone: "",
        address: "",
        address2: "",
        postcode: "",
        city: "",
        county: "",
        country: "",
      };
      this.selected = 0;
      document.getElementById("profileListId").value = 0;
    },
    async initProfileId() {
      let response = await axios.get("http://localhost:9000/profileId");
      IdProfileCount = response?.data;
      IdProfileCount++;
      console.log(IdProfileCount);
    },
  },
  async created() {
    this.initProfileId();
  },
};
</script>

<style scoped>
.add-form {
  margin-bottom: 30px;
}
.form-control {
  float: left;
  margin: 20px 100px;
  margin-top: 0px;
}
.form-control label {
  margin-top: 7px;
  display: block;
  color: cadetblue;
}
.form-control input {
  width: 300px;
  height: 35px;
  margin: 5px;
  padding: 3px 7px;
  font-size: 17px;
}
.form-control-check label {
  flex: 1;
  color: cadetblue;
}
.form-control-check input {
  flex: 2;
  height: 20px;
}
.drop {
  width: 300px;
  height: 34px;
  display: inline-block;
  background: white;
  color: black;
  border: none;
  padding: 3px 7px;
  margin: 5px;
  margin-top: 4px;
  border-radius: 1px;
  cursor: pointer;
  font-size: 17px;
  font-family: inherit;
}
.btn-blockProfile {
  display: block;
  width: 310px;
  height: 40px;
  margin-left: 100px;
  margin-right: 390px;
  margin-bottom: 15px;
}

.btn-blockMod {
  display: block;
  width: 310px;
  height: 40px;
  margin-left: 100px;
  margin-bottom: 15px;
}
.btn-del {
  margin-left: 200px;
}
.btnProfile {
  display: inline-block;
  background: #000;
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
}
.btnProfile:focus {
  outline: none;
}
.btnProfile:active {
  transform: scale(0.98);
}
.selectProfile {
  margin-top: 15px;
  margin-bottom: 15px;
}
.btnProfileDelete {
  display: inline-block;
  background: orangered;
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
}
.inner {
  display: inline-block;
}
</style>
