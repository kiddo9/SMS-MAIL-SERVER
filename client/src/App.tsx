import { BrowserRouter, Route, Routes } from "react-router-dom"
import Login from "./pages/auth/Login"
import NotFound from "./pages/NotFound"
import OTPVerify from "./pages/auth/OTPVerify"
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3"
import AuthContextProvider from "./contexts/AuthContext"
import Home from "./pages/Home"
import Footer from "./components/Footer"
import NavBar from "./components/NavBar"
import Templates from "./pages/Templates"

function App() {
  

  return (

    <BrowserRouter>
      <AuthContextProvider>
        
        <Routes>
          <Route path="/auth/*" element={
            <>
              <GoogleReCaptchaProvider reCaptchaKey={import.meta.env.VITE_RECAPTCHA_KEY}>
                <Routes>
                  <Route path="/login" element={<Login />} />
                  <Route path="/verify" element={<OTPVerify />} />
                  <Route path="*" element={<NotFound />} />
                </Routes>
              </GoogleReCaptchaProvider>
            </>
          } />
          <Route path="/*" element={
            <>
              <NavBar/>
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/templates" element={<Templates />} />
                <Route path="*" element={<NotFound />} />
              </Routes>
            </>
          } />
        </Routes>
        <Footer/>
      </AuthContextProvider>
    </BrowserRouter>
  )
}

export default App
