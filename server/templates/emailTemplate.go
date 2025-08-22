package templates

const OTPEmail = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Your One-Time Password</title>
  <style>
    body, p, h1, h2, h3, h4, h5, h6, table, td { margin:0; padding:0; font-family: Arial, sans-serif; }
    body { background-color: #f4f4f4; color: rgb(185,183,183); font-size: 16px; line-height: 1.5; }
    .email-container { max-width:600px; margin:0 auto; background-color:#252525; }
    .email-header { text-align:center; }
    .company-logo { max-height:95px; width:70%; object-fit:cover; }
    .email-body { padding:30px; }
    .otp-box { background-color:#6699ff; border:1px solid #e0e0e0; border-radius:8px; margin:24px 0; padding:20px; text-align:center; }
    .otp-code { font-size:32px; font-weight:bold; letter-spacing:4px; color:#f3e9e9; }
    .expiry-note { color:#fff; font-size:14px; margin-top:10px; }
    .divider { border:0; border-top:1px solid #e0e0e0; margin:24px 0; }
    .helpful-links { border-radius:8px; padding:20px; margin-top:24px; }
    .links-title { font-size:18px; margin-bottom:12px; }
    .link { color:#4a90e2; text-decoration:none; }
    .email-footer { padding:20px; text-align:center; font-size:12px; color:#777; }
  </style>
</head>
<body>
  <div class="email-container">
    <div class="email-header">
      <img src="https://techneo.ng/wp-content/uploads/2023/04/Artboard-1.png" alt="Company Logo" class="company-logo"/>
    </div>

    <div class="email-body">
      <p class="greeting">Hello {{.Name}},</p>
      <p>{{.EmailContent}}</p>

      <div class="otp-box">
        <p class="otp-code">{{.OTP}}</p>
        <p class="expiry-note">This code will expire in {{.ExpiryMinutes}} minutes.</p>
      </div>

      <hr class="divider" />

      <div class="helpful-links">
        <h3 class="links-title">Helpful Links</h3>
        <p>📚 <a href="" class="link">Help Center</a></p>
        <p>🔒 <a href="" class="link">Security Settings</a></p>
        <p>📱 <a href="" class="link">Download Our App</a></p>
        <p>📞 <a href="" class="link">Contact Us</a></p>
      </div>
    </div>

    <div class="email-footer">
      <div class="social-links">
        <a href="" class="social-link">Facebook</a> |
        <a href="" class="social-link">Twitter</a> |
        <a href="" class="social-link">Instagram</a> |
        <a href="" class="social-link">LinkedIn</a>
      </div>
      <p>© 2025 Neo cloud Technology. All rights reserved.</p>
    </div>
  </div>
</body>
</html>
`


const EmailTemplate1 = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Friendly Reminded</title>
  <style>
    body, p, h1, h2, h3, h4, h5, h6, table, td { margin:0; padding:0; font-family: Arial, sans-serif; }
    body { background-color: #f4f4f4; color: rgb(185,183,183); font-size: 16px; line-height: 1.5; }
    .email-container { max-width:600px; margin:0 auto; background-color:#252525; }
    .email-header { text-align:center; }
    .company-logo { max-height:95px; width:70%; object-fit:cover; }
    .email-body { padding:30px; }
    .otp-box { background-color:#6699ff; border:1px solid #e0e0e0; border-radius:8px; margin:24px 0; padding:20px; text-align:center; }
    .otp-code { font-size:32px; font-weight:bold; letter-spacing:4px; color:#f3e9e9; }
    .expiry-note { color:#fff; font-size:14px; margin-top:10px; }
    .divider { border:0; border-top:1px solid #e0e0e0; margin:24px 0; }
    .helpful-links { border-radius:8px; padding:20px; margin-top:24px; }
    .links-title { font-size:18px; margin-bottom:12px; }
    .link { color:#4a90e2; text-decoration:none; }
    .email-footer { padding:20px; text-align:center; font-size:12px; color:#777; }
  </style>
</head>
<body>
  <div class="email-container">
    <div class="email-header">
      <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRQaUVlv_0lrZg5J2v-eOUmRjEDP3v00c3gxw&s" alt="Company Logo" class="company-logo"/>
    </div>

    <div class="email-body">
      <p class="greeting">Hi {{.Name}},</p>
      <p>This is a friendly reminder about your outstanding balance of ₦{{.PendingPrice}} for your {{.course}} course. To ensure 
      continued access to your lectures, please settle your balance by or before {{.Date}}.<br/>
      For your convienience, you can make your payment through online bank transfer. <br/>

      if you've already made the payment, please meet with the admin with prove of payment Otherwise, you can contact the Admin department at {{.phoneNumber}} or 
      {{.EmailAddress}} to discuss payment options.

      Thanks,
      Neo cloud Admin Department.
      </p>

      <hr class="divider" />
    </div>

    <div class="email-footer">
      <div class="social-links">
        <a href="https://web.facebook.com/neocloudtech/" class="social-link">Facebook</a> |
        <a href="https://www.instagram.com/neocloudtech/" class="social-link">Instagram</a> |
        <a href="https://www.linkedin.com/company/neocloudtech" class="social-link">LinkedIn</a>
      </div>
      <p>© 2025 Neo cloud Technology. All rights reserved.</p>
    </div>
  </div>
</body>
</html>

`


const SmsTemp = `
Hi {{.Name}},

This is a friendly reminder about your outstanding balance of ₦{{.PendingPrice}} for your {{.course}} course. To ensure 
continued access to your lectures, please settle your balance by or before {{.Date}}.<br/>
For your convienience, you can make your payment through online bank transfer. <br/>

if you've already made the payment, please meet with the admin with prove of payment Otherwise, you can contact the Admin department at {{.phoneNumber}} or 
{{.EmailAddress}} to discuss payment options.

Thanks,
Neo cloud Admin Department.
`