Idea Code มั่วๆ
i. รับอินพุต 3 ตัวที่เป็นจำนวนเต็ม (bill, hour, minute)
ii. ตรวจสอบ minute >  0 ไหม
iii. ถ้ามากกว่า hour += 1
    แต่ถ้า ไม่มากกว่า -> ไม่ดำเนินการใดๆ
iv. ตรวจสอบ bill >= 1000 ไหม
v. ถ้ามากกว่า price = 30*(hour-4)
vi.ตรวจสอบ price < 0
vii. ถ้ามากกว่า price = 0
viii. แต่ถ้า bill < 1000 ไหม
ix. ถ้าน้อยกว่า price = 30*(hour-1)
            -> พิมพ์ price