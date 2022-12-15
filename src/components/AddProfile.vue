<template>
  <div>
    <select :key="i" v-for="(p, i) in profiles">
      <option :value="profile">{{ profile.fname }}</option>
    </select>
  </div>
  <form @submit="onSubmit" class="add-form">
    <div class="form-control">
      <label>First name</label>
      <input
        type="text"
        v-model="fname"
        name="Fname"
        placeholder="First name"
      />
    </div>
    <div class="form-control">
      <label>Last name</label>
      <input type="text" v-model="lname" name="Lname" placeholder="Last name" />
    </div>
    <div class="form-control">
      <label>Email</label>
      <input type="text" v-model="email" name="Email" placeholder="Email" />
    </div>
    <div class="form-control">
      <label>Phone </label>
      <input
        type="number"
        v-model="phone"
        name="Phone"
        placeholder="Phone number"
      />
    </div>
    <div class="form-control">
      <label>Address</label>
      <input type="text" v-model="address" name="Adress" placeholder="Adress" />
    </div>
    <div class="form-control">
      <label>Address2</label>
      <input
        type="text"
        v-model="address2"
        name="Address2"
        placeholder="Adress2"
      />
    </div>
    <div class="form-control">
      <label>Postcode</label>
      <input
        type="text"
        v-model="postcode"
        name="Postcode"
        placeholder="Postcode"
      />
    </div>
    <div class="form-control">
      <label>City</label>
      <input type="text" v-model="city" name="City" placeholder="City" />
    </div>
    <div class="form-control">
      <label>County</label>
      <input type="text" v-model="county" name="County" placeholder="County" />
    </div>
    <div class="form-control">
      <label>Country</label>
      <CountrySelect v-model="country" name="country" class="drop" />
    </div>
    <input
      type="submit"
      value="Save Profile"
      class="btnProfile btn-blockProfile ;"
    />
  </form>
</template>

<script>
import CountrySelect from "./CountrySelect.vue";
export default {
  name: "addProfile",
  components: {
    CountrySelect,
  },

  data() {
    return {
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
      props: {
        profiles: Array,
      },
    };
  },
  methods: {
    onSubmit(e) {
      e.preventDefault();
      console.log(this.profiles);
      if (!this.fname) {
        alert("Please add first name");
        return;
      }
      if (!this.lname) {
        alert("Please add last name");
        return;
      }
      if (!this.email) {
        alert("Please add email adress");
        return;
      }
      if (!this.phone) {
        alert("Please add phone number");
        return;
      }
      if (!this.address) {
        alert("Please add address");
        return;
      }
      if (!this.postcode) {
        alert("Please add postcode");
        return;
      }
      if (!this.city) {
        alert("Please add city");
        return;
      }
      if (!this.county) {
        alert("Please add county");
        return;
      }

      const newProfile = {
        lname: this.lanem,
        fname: this.fname,
        email: this.email,
        phone: this.phone,
        address: this.address,
        address2: this.address2,
        postcode: this.postcode,
        city: this.city,
        county: this.county,
        country: this.country,
      };

      this.$emit("add-profile", newProfile);

      this.fname = "";
      this.lname = "";
      this.email = "";
      this.phone = "";
      this.address = "";
      this.address2 = "";
      this.postcode = "";
      this.city = "";
      this.county = "";
      this.country = "";
    },
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
  margin-top: 5px;
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
</style>
