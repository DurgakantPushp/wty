<template>
<div>
<!-- Navbar (sit on top) -->
<div class="w3-top">
  <div style = "color:#fff!important;background-color:rgb(6, 110, 128)!important; padding:16px 16px" class="w3-bar w3-white w3-card" id="myNavbar">
    <a href="#home" class="w3-bar-item w3-button w3-wide" style="margin-top: 0px; padding:0px 0px">
      <img src = "../assets/wty_white_logo-home.png">
    </a>
    <!-- Right-sided navbar links -->
    <div class="w3-right w3-hide-small">
        <button 
          class="w3-bar-item w3-button"
          data-toggle="modal"
          data-target="#loginModal"
        >
          <i class="fa fa-user"></i>
          LOGIN
        </button>
        <a href="#stories" class="w3-bar-item w3-button">
          <i class="fa fa-th"></i>
          GRATITUDE STORIES
        </a>
        <button 
          class="w3-bar-item w3-button"
          id="myBtn3"
          type="button"
        >
          <i class="fa fa-paper-plane"></i>
          SEND GRATITUDE
        </button>
    </div>
    <!-- Hide right-floated links on small screens and replace them with a menu icon -->

    <a href="javascript:void(0)" class="w3-bar-item w3-button w3-right w3-hide-large w3-hide-medium" onclick="w3_open()">
      <i class="fa fa-bars"></i>
    </a>
  </div>
</div>

<!-- Sidebar on small screens when clicking the menu icon -->
<nav class="w3-sidebar w3-bar-block w3-black w3-card w3-animate-left w3-hide-medium w3-hide-large" style="display:none" id="mySidebar">
  <a href="javascript:void(0)" onclick="w3_close()" class="w3-bar-item w3-button w3-large w3-padding-16">Close ×</a>
  <button class="w3-bar-item w3-button" id="btnLogin" type="button"><i class="fa fa-user"></i> LOGIN</button>
  <a href="#stories" class="w3-bar-item w3-button"><i class="fa fa-th"></i> GRATITUDE STORIES</a>
  <button 
    class="w3-bar-item w3-button"
    id="btnSendGrat" type="button">
    <i class="fa fa-paper-plane"></i>
    SEND GRATITUDE
  </button>  
</nav>

<!-- Header with full-height image -->
<header style = "background-color:rgb(0, 0, 0)!important; " class="bgimg-1 w3-display-container w3-grayscale-min" id="home">
  <div class="w3-display-left w3-text-white" style="padding:48px;">
    <span class="w3-jumbo w3-hide-small">Let's make gratitude a lifestyle!</span><br>
    <span class="w3-xxlarge w3-hide-large w3-hide-medium">Let's make gratitude <br>a lifestyle!</span><br>
    <span class="w3-large">WannaThankYou is new and unique social app<br> that helps you express
      gratitude and<br> show your appreciation. </span>
    <button 
      class="w3-button w3-padding-large w3-large w3-margin-top w3-opacity w3-hover-opacity-off btn btn-primary"
      id="myBtn1"
    >
      Register here and start today
    </button>
  </div> 
  <div class="w3-display-bottomleft w3-text-grey w3-large" style="padding:24px 48px">
    <i class="fa fa-facebook-official w3-hover-opacity"></i>
    <i class="fa fa-instagram w3-hover-opacity"></i>
    <i class="fa fa-snapchat w3-hover-opacity"></i>
    <i class="fa fa-pinterest-p w3-hover-opacity"></i>
    <i class="fa fa-twitter w3-hover-opacity"></i>
    <i class="fa fa-linkedin w3-hover-opacity"></i>
  </div>
</header>

<!-- Share Your Stories -->
<div class="w3-container" style="padding:128px 16px" id="stories">
  <h3 class="w3-center">Gratitude Stories</h3>
  <p class="w3-center w3-large">Share your stories and be inspired by others.</p>
  <div class="w3-row-padding w3-grayscale" style="margin-top:64px">
    <div class="w3-col l3 m6 w3-margin-bottom">
      <div class="w3-card">
        <img src="../assets/img_avatar1.png" alt="John" style="width:100%">
        <div class="w3-container">
          <h3>John</h3>
          <p class="w3-opacity">Artist</p>
          <p>Shortly after seeing The Secret movie, I was at a wedding for a friend and they had gratitude rocks on the table as a gift. Now I knew from watching the movie what these were used for, so I began a process of feeling grateful every day for the <a href="#">Read more....</a></p>
        </div>
      </div>
    </div>
    <div class="w3-col l3 m6 w3-margin-bottom">
      <div class="w3-card">
        <img src="../assets/img_avatar2.png" alt="John" style="width:100%">
        <div class="w3-container">
          <h3>Anja Doe</h3>
          <p class="w3-opacity">Doctor</p>
          <p>Shortly after seeing The Secret movie, I was at a wedding for a friend and they had gratitude rocks on the table as a gift. Now I knew from watching the movie what these were used for, so I began a process of feeling grateful every day for the <a href="#">Read more....</a></p>
        </div>
      </div>
    </div>
    <div class="w3-col l3 m6 w3-margin-bottom">
      <div class="w3-card">
        <img src="../assets/img_avatar3.jpeg" alt="John" style="width:100%">
        <div class="w3-container">
          <h3>Mike Ross</h3>
          <p class="w3-opacity">Gardener</p>
          <p>Shortly after seeing The Secret movie, I was at a wedding for a friend and they had gratitude rocks on the table as a gift. Now I knew from watching the movie what these were used for, so I began a process of feeling grateful every day for the <a href="#">Read more....</a></p>
        </div>
      </div>
    </div>
    <div class="w3-col l3 m6 w3-margin-bottom">
      <div class="w3-card">
        <img src="../assets/img_avatar4.jpeg" alt="Dan" style="width:100%">
        <div class="w3-container">
          <h3>Dan Star</h3>
          <p class="w3-opacity">Cricketer</p>
          <p>Shortly after seeing The Secret movie, I was at a wedding for a friend and they had gratitude rocks on the table as a gift. Now I knew from watching the movie what these were used for, so I began a process of feeling grateful every day for the <a href="#">Read more....</a></p>
        </div>
      </div>
    </div>
  </div>
</div>

<!-- Promo Section "Statistics" -->
<div class="w3-container w3-row w3-center w3-dark-grey w3-padding-64">
  <div class="w3-quarter">
    <span class="w3-xxlarge">14k+</span>
    <br>WTY Team Members
  </div>
  <div class="w3-quarter">
    <span class="w3-xxlarge">55k+</span>
    <br>Entities
  </div>
  <div class="w3-quarter">
    <span class="w3-xxlarge">89k+</span>
    <br>Business Users
  </div>
  <div class="w3-quarter">
    <span class="w3-xxlarge">150k+</span>
    <br>Users
  </div>
</div>

<!-- Modal for full size images on click-->
<div id="modal01" class="w3-modal w3-black" onclick="this.style.display='none'">
  <span class="w3-button w3-xxlarge w3-black w3-padding-large w3-display-topright" title="Close Modal Image">×</span>
  <div class="w3-modal-content w3-animate-zoom w3-center w3-transparent w3-padding-64">
    <img id="img01" class="w3-image">
    <p id="caption" class="w3-opacity w3-large"></p>
  </div>
</div>

<!-- Contact Section -->
<div class="w3-container w3-light-grey" style="padding:128px 16px" id="contact">
  <h3 class="w3-center">CONTACT</h3>
  <p class="w3-center w3-large">Please let us know what we can do better. We want to hear from you.</p>
  <div class="w3-row-padding" style="margin-top:64px">
    <div class="w3-half">
      <p><i class="fa fa-map-marker fa-fw w3-xxlarge w3-margin-right"></i> Bangalore, India</p>
      <p><i class="fa fa-phone fa-fw w3-xxlarge w3-margin-right"></i> Phone: +XXX-XXXXXX</p>
      <p><i class="fa fa-envelope fa-fw w3-xxlarge w3-margin-right"> </i> Email: mail@mail.com</p>
      <br>
      <form action="/action_page.php" target="_blank">
        <p><input class="w3-input w3-border" type="text" placeholder="Name" required name="Name"></p>
        <p><input class="w3-input w3-border" type="text" placeholder="Email" required name="Email"></p>
        <p><input class="w3-input w3-border" type="text" placeholder="Subject" required name="Subject"></p>
        <p><input class="w3-input w3-border" type="text" placeholder="Message" required name="Message"></p>
        <p>
          <button class="w3-button w3-black" type="submit">
            <i class="fa fa-paper-plane"></i> SEND MESSAGE
          </button>
        </p>
      </form>
    </div>
    <div class="w3-half">
      <!-- Add Google Maps -->
      <div id="googleMap" class="w3-greyscale-max" style="width:100%;height:510px;"></div>
    </div>
  </div>
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
 
<!-- Login -->
<login/>

<!-- Register -->
<register/>

<!-- Gratitude -->
<gratitude/>

</div>
</template>

<script>
/* eslint-disable */

import Register from './Register'
import Login from './Login.vue'
import Gratitude from './Gratitude'

export default {
  components: {
    'register': Register,
    'login': Login,
    'gratitude': Gratitude
  }
}

$(document).ready(function(){
    $("#btnLogin").click(function(){
        console.log('login button')
        $("#loginModal").modal();
    });
});
$(document).ready(function(){
    $("#myBtn2").click(function(){
        $("#myModal").modal();
    });
});
$(document).ready(function(){
    $("#myBtn1").click(function(){
        $("#modalReg").modal()
    });
});
$(document).ready(function(){
    $("#myBtn3").click(function(){
        $("#modalGrat").modal();
    });
});
$(document).ready(function(){
    $("#btnSendGrat").click(function(){
        $("#modalGrat").modal();
    });
});

//Add Google Maps
function myMap()
{
  myCenter=new google.maps.LatLng(41.878114, -87.629798);
  var mapOptions= {
    center:myCenter,
    zoom:12, scrollwheel: true, draggable: true,
    mapTypeId:google.maps.MapTypeId.ROADMAP
  };
  var map=new google.maps.Map(document.getElementById("googleMap"),mapOptions);

  var marker = new google.maps.Marker({
    position: myCenter,
  });
  marker.setMap(map);
}

// Toggle between showing and hiding the sidebar when clicking the menu icon
var mySidebar = document.getElementById("mySidebar");

function w3_open() {
    if (mySidebar.style.display === 'block') {
        mySidebar.style.display = 'none';
    } else {
        mySidebar.style.display = 'block';
    }
}

// Close the sidebar with the close button
function w3_close() {
    mySidebar.style.display = "none";
}
</script>

<style scoped>
  .modal-header, .close {
      background-color: #5cb85c;
      color:white !important;
      text-align: center;
      font-size: 30px;
  }
  .modal-footer {
      background-color: #f9f9f9;
  }
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
