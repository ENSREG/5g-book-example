# EN 帶你入門 5G 核心網路 程式範例

> 相關演講：
> - [[OSN Day 2023] How to start contributing to free5GC](https://youtu.be/OQEMNJ-7rJs?si=XF6cHCRcIz3aJ7RX) [[PPT](https://github.com/ENSREG/5g-book-example/blob/main/yi%20chen.pptx)]
> - [[COSCUP 2022] EN 帶你入門 5G 核心網路](https://www.youtube.com/watch?v=plFFX_geCZs)

![image](https://github.com/ENSREG/5g-book-example/assets/42661015/2daaa1dd-b6ce-4749-bffd-020c18e9c116)


## 勘誤/補充表

感謝以下同學協助勘誤：
- [劉又聖](https://www.linkedin.com/in/yu-sheng-liu-41a45a24a/)
- [許博勝](https://www.linkedin.com/in/%E5%8D%9A%E5%8B%9D-%E8%A8%B1-392a87205/)

### p1-75 configuration update

Configuration update 應該是由核心網路端發起：
- AMF 發送 Configuration Update Command
- UE 回應 Configuration Update Complete

### p1-85 QoS Rule

書中描述：
> 如果 UE 沒有找到合適的 PDR，那麼 UE 會丟棄這個 UL 封包。

這裡的 PDR 應改為 QoS Rule，PDR 是存放在 5G Core 端的 UPF 內，而非存放於 UE 端。

### p1-88 RFSP

RFST 應該為 RFSP。
