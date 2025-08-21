import { useEffect, useState } from "react";
import AdminClient from "../../lib/adminClient"
import {OtpRequest, OtpVerificationRequest} from "../../proto/Admin"
import { useNavigate, useSearchParams } from "react-router-dom";
import { useAuthContext } from "../../contexts/AuthContext";

const OTPVerify = () => {
    const [otp, setOtp] = useState('');
    const [allowReset, setAllowReset] = useState(false);
    const [email, setEmail] = useState('');
    const [error, setError] = useState('');
    const [alert, setAlert] = useState('');
    const [loading, setLoading] = useState(false);
    const [searchParams] = useSearchParams()
    const nav = useNavigate()
    const {setAtkFunc} = useAuthContext();

    const tk = searchParams.get('tk');
    
    useEffect(() => {
        if(!tk) return
        AdminClient.validateToken({token: tk}).then((response) => {
            if(!response || !response.response.isValid) return
            setEmail(response.response.email);
        })
    }, [tk])
    // const [password, setPassword] = useState('');
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        try {
            e.preventDefault();

            const response = await AdminClient.verifyOtp(OtpVerificationRequest.create({email: email, otp: otp}),{})

            if(!response || !response.response.isVerified){
                setError("Failed Verification");
                setAlert('');
                return;
            }


            setAtkFunc(response.response.message);
            nav('/');
            
        } catch (error) {
            if(import.meta.env.VITE_ENV === "development") console.log(error);
            setError("Failed to login");
            setAlert('');
        }finally{
            setLoading(false);
        }
        
    }

    const handleOTPResend = async() =>{
        try {
            setAllowReset(false);
            const request = await AdminClient.sendOtp(OtpRequest.create({email: email}),{})

            if(!request || !request.response.message){
                setError("Failed to resend OTP");
                setAlert('');
                return;
            }
            setAlert("OTP resent successfully");
            setError('');
        } catch (error) {
            if(import.meta.env.VITE_ENV === "development") console.log(error);
            setError("Failed to login");
            setAlert('');
        }finally{
            setLoading(false);
        }
        
    }

    useEffect(() => {
        if(!allowReset){
            const timer = setTimeout(() => {
                setAllowReset(true);
            }, 5000);
            return () => clearTimeout(timer);
        }
        else return
    }, [allowReset])

    if(!tk){ 
        setTimeout(() => {
            nav('/auth/login');
        }, 2000)
        return <div className='flex flex-col h-screen justify-center font-semibold items-center bg-[#6699ff]/10'>Invalid Request. Redirecting...</div>
    }
    if(!email) return <div className='flex flex-col h-screen justify-center font-semibold items-center bg-[#6699ff]/10'>Loading...</div>
  return (
    <div className='flex flex-col h-screen justify-center items-center bg-[#6699ff]/10'>
        
        <form onSubmit={handleSubmit} className='bg-white p-4 rounded-md shadow-md flex flex-col min-w-[300px] max-w-[400px]'>
            <div className="flex flex-row-reverse  items-center justify-center">
                <img className='' width={30} src="/logo-icon.png" alt="Logo" />
                <h1 className='text-lg font-bold mb-1 text-center'>Verify OTP</h1>
                <img className='rotate-y-180' width={30} src="/logo-icon.png" alt="Logo" />
            </div>
            <span className='block mb-1 text-xs text-center text-gray-500 '>An OTP has been sent to your email</span>
            <div className="mt-2 mb-1">
                
                <input value={otp} onChange={(e) => setOtp(e.target.value)} className='p-2 border-2 focus:border-[#6699ff] border-gray-200 rounded-md w-full' type="text" placeholder='Enter the OTP...' />
            </div>
            <button disabled={loading} type="submit" className='mt-1 bg-[#6699ff] hover:bg-[#6699ff]/80 cursor-pointer text-white p-2 rounded-md'>{loading ? "Verifying..." : "Verify"}</button>
            {error && <span className='mt-1 text-xs text-center text-red-500'>{error}</span>}
            {alert && <span className='mt-1 text-xs text-center text-green-500'>{alert}</span>}
            <button disabled={!allowReset} onClick={handleOTPResend} className='mt-1 bg-transparent self-start text-xs font-semibold cursor-pointer text-[#6699ff] disabled:text-gray-500 disabled:cursor-not-allowed'>Resend</button>
        </form>
    </div>

  )
}

export default OTPVerify