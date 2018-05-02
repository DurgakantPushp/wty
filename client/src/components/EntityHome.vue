<template>
<div id="entityHome">
  <!-- Navbar (sit on top) -->
  <div class="w3-top">
    <div style = "color:#fff!important;background-color:rgb(6, 110, 128)!important; padding:16px 16px" class="w3-bar w3-white w3-card" id="myNavbar">
      <a href="#home" class="w3-bar-item w3-button w3-wide"
        style="margin-top: 0px; padding:0px 0px">
        <img src = "../assets/wty_white_logo-home.png">
      </a>
      <!-- Right-sided navbar links -->
      <div class="w3-right w3-hide-small">
        <a href="#recent" class="w3-bar-item w3-button"><i class="glyphicon glyphicon-eye-open"></i> VIEW RECENT GRATITUDE</a>
        <a href="#received" class="w3-bar-item w3-button"><i class="glyphicon glyphicon-eye-open"></i>VIEW RECEIVED GRATITUDE</a>
        <a href="#sent" class="w3-bar-item w3-button"><i class="glyphicon glyphicon-eye-open"></i>VIEW SENT GRATITUDE</a>
        <button class="w3-bar-item w3-button" id="myBtn3" type="button"><i class="fa fa-paper-plane"></i> SEND GRATITUDE</button>
      </div>
      <!-- Hide right-floated links on small screens and replace them with a menu icon -->

      <a href="javascript:void(0)"
        class="w3-bar-item w3-button w3-right w3-hide-large w3-hide-medium"
        @click="w3_open">
        <i class="fa fa-bars"></i>
      </a>
    </div>
  </div>

  <!-- Sidebar on small screens when clicking the menu icon -->
          <nav class="w3-sidebar w3-bar-block w3-black w3-card w3-animate-left w3-hide-medium w3-hide-large" style="display:none" id="mySidebar">
            <a href="#recent" class="w3-bar-item w3-button"><i class="glyphicon glyphicon-eye-open"></i> VIEW RECENT GRATITUDE</a>
            <a href="#received" class="w3-bar-item w3-button"><i class="glyphicon glyphicon-eye-open"></i>VIEW RECEIVED GRATITUDE</a>
            <a href="#sent" class="w3-bar-item w3-button"><i class="glyphicon glyphicon-eye-open"></i>VIEW SENT GRATITUDE</a>
            <button class="w3-bar-item w3-button" id="myBtn2" type="button"><i class="fa fa-paper-plane"></i> SEND GRATITUDE</button>
          </nav>

  <!-- Header with full-height image -->
  <header class="bgimg-1 w3-display-container w3-grayscale-min" id="home">
          <div class="w3-display-left w3-text-white" style="padding:48px">
            <span class="w3-jumbo w3-hide-small">Start something that matters</span><br>
            <span class="w3-xxlarge w3-hide-large w3-hide-medium">Start something that matters</span><br>
            <span class="w3-large">WannaThankYou is new and unique social app<br> that helps you express
              gratitude and<br> show your appreciation.</span>
            <p><a href="#about" class="w3-button w3-white w3-padding-large w3-large w3-margin-top w3-opacity w3-hover-opacity-off">Manage Your Profile</a></p>
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

  <!-- Recent Gratitude Section -->
  <div class="w3-container bgimg-1" style="padding:128px 16px" id="recent">
          <h3 class="w3-center">Recent Gratitude</h3>
          <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
            <div class="w3-container" style="width:100%">
              <h3>Jogdf Hdfdge</h3>
              <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
              <p>Phasellus eget enim eu lectus faucibus gmy
                vestibulum. Suspendisse sodales pellentesque elementum.
                vestibulum. Suspendisse sodales pellentesque elementum.
                vestibulum. Suspendisse sodales pellentesque elementum.
                vestibulum. Suspendisse sodales pellentesque elementum.</p>
            </div>
          </div>

          <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
              <div class="w3-container" style="width:100%">
                <h3>John Doe</h3>
                <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
                <p>Phasellus eget enim eu lectus faucibus gmy
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.</p>
              </div>
          </div>

          <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
              <div class="w3-container" style="width:100%">
                <h3>John Doe</h3>
                <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
                <p>Phasellus eget enim eu lectus faucibus gmy
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.</p>
              </div>
          </div>

          <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
              <div class="w3-container" style="width:100%">
                <h3>John Doe</h3>
                <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
                <p>Phasellus eget enim eu lectus faucibus gmy
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.
                  vestibulum. Suspendisse sodales pellentesque elementum.</p>
              </div>
          </div>

  </div>

  <!-- Received Gratitude Section -->
  <div class="w3-container bgimg-1" style="padding:128px 16px" id="received">
      <h3 class="w3-center">Received Gratitude</h3>
      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
        <div class="w3-container" style="width:100%">
          <h3>ABC Edu Trust</h3>
          <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
          <p>Phasellus eget enim eu lectus faucibus gmy
            vestibulum. Suspendisse nlkmjgthrt io  ito jirthj  hjroijhrlhjrjprosodales pellentesque elementum.
            vestibulum. Suspendisse sodales pellentesque elementum.
            vestibulum. Suspendisse sodales pellentesque elementum.
            vestibulum. Suspendisse sodales pellentesque elementum.</p>
        </div>
      </div>

      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
          <div class="w3-container" style="width:100%">
            <h3>John Doe</h3>
            <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
            <p>Phasellus eget enim eu lectus faucibus gmy
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.</p>
          </div>
      </div>

      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
          <div class="w3-container" style="width:100%">
            <h3>John Doe</h3>
            <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
            <p>Phasellus eget enim eu lectus faucibus gmy
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.</p>
          </div>
      </div>

      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
          <div class="w3-container" style="width:100%">
            <h3>John Doe</h3>
            <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
            <p>Phasellus eget enim eu lectus faucibus gmy
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.</p>
          </div>
      </div>

  </div>

  <!-- Sent Gratitude Section -->
  <div class="w3-container bgimg-1" style="padding:128px 16px" id="sent">
      <h3 class="w3-center">Sent Gratitude</h3>
      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
        <div class="w3-container" style="width:100%">
          <h3>ABC Edu Trust</h3>
          <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
          <p>Phasellus eget enim eu lectus faucibus gmy
            vestibulum. Suspendisse nlkmjgthrt io  ito jirthj  hjroijhrlhjrjprosodales pellentesque elementum.
            vestibulum. Suspendisse sodales pellentesque elementum.
            vestibulum. Suspendisse sodales pellentesque elementum.
            vestibulum. Suspendisse sodales pellentesque elementum.</p>
        </div>
      </div>

      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
          <div class="w3-container" style="width:100%">
            <h3>John Doe</h3>
            <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
            <p>Phasellus eget enim eu lectus faucibus gmy
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.</p>
          </div>
      </div>

      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
          <div class="w3-container" style="width:100%">
            <h3>John Doe</h3>
            <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
            <p>Phasellus eget enim eu lectus faucibus gmy
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.</p>
          </div>
      </div>

      <div class="w3-card" style="margin-right:200px;margin-left:200px; background-color:rgba(250, 246, 246, 0.432);">
          <div class="w3-container" style="width:100%">
            <h3>John Doe</h3>
            <p class="w3-opacity" Style="font-size:12px; color:green;">12 March, 2018,</p>
            <p>Phasellus eget enim eu lectus faucibus gmy
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.
              vestibulum. Suspendisse sodales pellentesque elementum.</p>
          </div>
      </div>

  </div>

  <!-- Footer -->
  <footer class="w3-center w3-black w3-padding-64">
    <a href="#home" class="w3-button w3-light-grey"><i class="fa fa-arrow-up w3-margin-right"></i>To the top</a>
    <div class="w3-xlarge w3-section">
      <i class="fa fa-facebook-official w3-hover-opacity"></i>
      <i class="fa fa-instagram w3-hover-opacity"></i>
      <i class="fa fa-snapchat w3-hover-opacity"></i>
      <i class="fa fa-pinterest-p w3-hover-opacity"></i>
      <i class="fa fa-twitter w3-hover-opacity"></i>
      <i class="fa fa-linkedin w3-hover-opacity"></i>
    </div>
  </footer>

  <!-- Gratitude Message -->
  <!-- Modal -->
  <div class="modal fade" id="myModal2" role="dialog">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header" style="padding:35px 50px; background-color:rgb(6, 110, 128)!important">
          <button type="button" class="close" data-dismiss="modal">&times;</button>
          <h4><img src = "wty_white_logo-home.png"></h4>
        </div>
        <div class="modal-body" style="padding:40px 50px;">
          <form role="form">
            <div class="form-group">
              <label for="usr"><span class="glyphicon glyphicon-user"></span> Receiver's Name</label>
              <input type="text" class="form-control" id="usr" placeholder="Enter receiver's email/User ID/Profile Name">
            </div>
            <div class="form-group">
              <label for="msg"><span class="glyphicon fa fa-envelope"></span> Your Gratitude</label>
              <input type="text" class="form-control" id="msg" placeholder="Type your message">
            </div>
              <button type="submit" class="btn btn-success btn-block" style="background-color:rgb(6, 110, 128)!important"><span class="glyphicon glyphicon-off"></span> Send Gratitude</button>
          </form>
        </div>
        <div class="modal-header" style="padding:35px 50px; background-color:rgb(6, 110, 128)!important">
            <h4><img src = "wty_white_logo-home.png"></h4>
        </div>
      </div>

    </div>
  </div>
</div>
</template>

<script>
export default {
  data () {
    return {
      name: ''
    }
  },
  methods: {
    w3_open () {
      var mySidebar = document.getElementById("mySidebar")
      if (mySidebar.style.display === 'block') {
          mySidebar.style.display = 'none';
      } else {
          mySidebar.style.display = 'block';
      }
    },
    w3_close () {
      var mySidebar = document.getElementById("mySidebar")
      mySidebar.style.display = "none";
    }
  }
  // $(document).ready(function(){
  //   $("#myBtn3").click(function(){
  //       $("#myModal2").modal();
  //   });
  // });
  // $(document).ready(function(){
  //     $("#myBtn2").click(function(){
  //         $("#myModal2").modal();
  //     });
  // });
 // Toggle between showing and hiding the sidebar when clicking the menu icon
}
</script>

<style scoped>
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

