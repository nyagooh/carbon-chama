class LoginForm extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    const container = document.createElement("div");
    container.classList.add("login-container");
    const link = document.createElement("link");
    link.rel = "stylesheet";
    link.type = "text/css";
    link.href = "/static/login.css";

    container.innerHTML = `
        <div class="left-section">
             <h2>Capturing Moments, Creating Memories</h2>
                    </div>
         <div class="right-section">
                        <h2>Welcome Back</h2>
                        <p>Please Enter your Account details</p>
                        <form id="login-form" action="/auth/login" method="post">
                            <input type="email" name="email" placeholder="Email" required>
                            <input type="password" name="password" placeholder="Enter your password" required>
                           <div class="terms">
                                <a href="#">Forgot Password?</a>
                            </div>
                            <button type="submit">Login</button>
                        </form>
                        <p>Or register with</p>
                        <button id="login-btn" class="social-btn">
                        <text class = "google-text" x="10" y="30" font-family="Poppins, Arial, sans-serif"  font-weight="bold">
                                <span class="blue">G</span>
                                <span class="red">o</span>
                                <span class="yellow">o</span>
                                <span class="blue">g</span>
                                <span class="green">l</span>
                                <span class="red">e</span>
                        </text>
                        </button>
                     <div class="create-account">
                  <p>Don't have an account? <a href="/signup">Create an account</a></p>
                </div>
         </div>

        `;

    this.shadowRoot.append(link, container);
    this.shadowRoot.querySelector("#login-btn")
      .addEventListener("click", () => {
        window.location.href = "/auth/google";
      });
      
    // Add event listener for form submission
    this.shadowRoot.querySelector("#login-form")
      .addEventListener("submit", (e) => {
        // Form will submit normally to the action URL
      });
  }
}
customElements.define("login-form", LoginForm);
