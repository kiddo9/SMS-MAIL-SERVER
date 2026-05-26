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
import SMSBalances from "./pages/SMSBalances"
import { ToastContainer } from "react-toastify"

function App() {
  

  return (

    <BrowserRouter>
    <GoogleReCaptchaProvider reCaptchaKey={import.meta.env.VITE_RECAPTCHA_KEY}>
      <AuthContextProvider>
      <ToastContainer position="top-right" autoClose={5000} hideProgressBar={false} newestOnTop={false} closeOnClick rtl={false} pauseOnFocusLoss draggable pauseOnHover />
        <Routes>
          <Route path="/auth/*" element={
            <>
              
                <Routes>
                  <Route path="/login" element={<Login />} />
                  <Route path="/verify" element={<OTPVerify />} />
                  <Route path="*" element={<NotFound />} />
                </Routes>
              
            </>
          } />
          <Route path="/*" element={
            <>
              <NavBar/>
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/templates" element={<Templates />} />
                <Route path="/balances" element={<SMSBalances/>} />
                <Route path="*" element={<NotFound />} />
              </Routes>
            </>
          } />
        </Routes>
        <Footer/>
      </AuthContextProvider>
      </GoogleReCaptchaProvider>
    </BrowserRouter>
  )
}

export default App
