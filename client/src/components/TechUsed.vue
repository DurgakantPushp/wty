<template>
  <div>
    <!-- Navbar (sit on top) -->
    <div class="w3-top">
      <div style = "color:#fff!important;background-color:rgb(6, 110, 128)!important; padding:16px 16px" class="w3-bar w3-white w3-card" id="myNavbar">
        <a href="#home" class="w3-bar-item w3-button w3-wide" style="margin-top: 0px; padding:0px 0px">
          <img src = "wty_white_logo-home.png">
        </a>
        <!-- Right-sided navbar links -->
        <div class="w3-right w3-hide-small">
          <a href="#about" class="w3-bar-item w3-button">TECHNOLOGY USED</a>
          <!-- Hide right-floated links on small screens and replace them with a menu icon -->

          <a href="javascript:void(0)" class="w3-bar-item w3-button w3-right w3-hide-large w3-hide-medium" onclick="w3_open()">
            <i class="fa fa-bars"></i>
          </a>
        </div>
      </div>
    </div>

    <!-- Sidebar on small screens when clicking the menu icon -->
    <nav class="w3-sidebar w3-bar-block w3-black w3-card w3-animate-left w3-hide-medium w3-hide-large" style="display:none" id="mySidebar">
      <a href="javascript:void(0)" onclick="w3_close()" class="w3-bar-item w3-button w3-large w3-padding-16">Close ×</a>
      <a href="#about" onclick="w3_close()" class="w3-bar-item w3-button">TECHNOLOGY USED</a>
    </nav>

    <!-- Used Technology Description Section -->
    <div class="w3-container" style="padding:128px 16px" id="about">
            <h3 class="w3-center">ABOUT THE FRONTEND TECHNOLOGY INVOLVED</h3>

            <div class="w3-row-padding w3-center" style="margin-top:64px">
              <div class="w3-quarter">
                    <img src = "vuejs.png" >
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore.</p>
              </div>
              <div class="w3-quarter">
                    <img src = "bootstrap.png" >
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore.</p>
              </div>
              <div class="w3-quarter">
                <i class="fa fa-diamond w3-margin-bottom w3-jumbo"></i>
                <p class="w3-large">Design</p>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore.</p>
              </div>
              <div class="w3-quarter">
                <i class="fa fa-cog w3-margin-bottom w3-jumbo"></i>
                <p class="w3-large">Support</p>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore.</p>
              </div>
            </div>
            <p class="w3-center w3-large">
              It is SPA. Frontend is decoupled from backend. It is component based. vuejs is main framework used to build this. Other technology used
              webpack to compile frontend, yarn to manage packages, vue-router to handle frontend routing, vuex to manage component state, axios to handle http request
              Bootstrap to beautify pages, swal for alert, error, success pop ups
            </p>
    </div>

    <div class="w3-container" style="padding:128px 16px" id="about">
      <h3 class="w3-center">BACKEND TECHNOLOGY INVOLVED</h3>
      <div class="w3-row-padding w3-center" style="margin-top:64px">

        <div class="w3-quarter">
          <i class="fa fa-diamond w3-margin-bottom w3-jumbo"></i>
          <p class="w3-large">Design</p>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore.</p>
        </div>
        <div class="w3-quarter">
          <i class="fa fa-cog w3-margin-bottom w3-jumbo"></i>
          <p class="w3-large">Support</p>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore.</p>
        </div>
      </div>
      <p class="w3-center w3-large">It is SPA. Frontend is decoupled from backend. It is component based. vuejs is main framework used to build this. Other technology used
        webpack to compile frontend, yarn to manage packages, vue-router to handle frontend routing, vuex to manage component state, axios to handle http request
        Bootstrap to beautify pages, swal for alert, error, success pop ups
      </p>
    </div>

    <!-- Footer -->
    <footer class="w3-center w3-black w3-padding-64">
      <a href="#home" class="w3-button w3-light-grey">
        <i class="fa fa-arrow-up w3-margin-right"></i>
        To the top
      </a>
      <div class="w3-xlarge w3-section">
        <i class="fa fa-facebook-official w3-hover-opacity"></i>
        <i class="fa fa-instagram w3-hover-opacity"></i>
        <i class="fa fa-snapchat w3-hover-opacity"></i>
        <i class="fa fa-pinterest-p w3-hover-opacity"></i>
        <i class="fa fa-twitter w3-hover-opacity"></i>
        <i class="fa fa-linkedin w3-hover-opacity"></i>
      </div>
    </footer>
  </div>
</template>

<script>
// // Toggle between showing and hiding the sidebar when clicking the menu icon
// var mySidebar = document.getElementById("mySidebar");

// function w3_open() {
//     if (mySidebar.style.display === 'block') {
//         mySidebar.style.display = 'none';
//     } else {
//         mySidebar.style.display = 'block';
//     }
// }

// // Close the sidebar with the close button
// function w3_close() {
//     mySidebar.style.display = "none";
// }
export default {
  data () {
    return {
      name: 'ajay'
    }
  }
}
</script>

 <style>
  body,h1,h2,h3,h4,h5,h6 {font-family: "Raleway", sans-serif}
  body, html {
    height: 100%;
    line-height: 1.8;
  }
  /* Full height image header */
  .bgimg-1 {
    background-position: center;
    background-size: cover;
    background-image: url("../assets/wty-landing-page-hero-image.jpg");
    min-height: 100%;
  }
  .w3-bar .w3-button {
    padding-top: 16px;
  }
</style>
