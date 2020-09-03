package lang

// LDA Load Accumulator with Memory
const LDA = `LDA`

// LDX Load Index X with Memory
const LDX = `LDX`

// LDY Load Index Y with Memory
const LDY = `LDY`

// STA Store Accumulator in Memory
const STA = `STA`

// STX Store Index X in Memory
const STX = `STX`

// STY Store Index Y in Memory
const STY = `STY`

// ADC Add Memory to Accumulator with Carry
const ADC = `ADC`

// SBC Subtract Memory from Accumulator with Borrow
const SBC = `SBC`

// INC Increment Memory by One
const INC = `INC`

// INX Incrememnt Index X by One
const INX = `INX`

// INY Increment Index Y by One
const INY = `INY`

// DEC Decrement Memory by One
const DEC = `DEC`

// DEX Decrement Index X by One
const DEX = `DEX`

// DEY Decrement Index Y by One
const DEY = `DEY`

// ASL Arithmetic Shift Left One Bit
const ASL = `ASL`

// LSR Logical Shift Right One Bit
const LSR = `LSR`

// ROL Rotate Left One Bit
const ROL = `ROL`

// ROR Rotate Right One Bit
const ROR = `ROR`

// AND AND Memory with Accumulator
const AND = `AND`

// ORA OR Memory with Accumulator
const ORA = `ORA`

// XOR XOR Memory with Accumulator
const XOR = `XOR`

// CMP Compare Memory and Accumulator
const CMP = `CMP`

// CPX Compare Memory and Index X
const CPX = `CPX`

// CPY Compare Memory with Index Y
const CPY = `CPY`

// BIT Test Bits in Memory with Accumulator
const BIT = `BIT`

// BCC Branch on Carry Clear
const BCC = `BCC`

// BCS Branch on Carry Set
const BCS = `BCS`

// BNE Branch on Result not Zero
const BNE = `BNE`

// BEQ Branch on Result Zero
const BEQ = `BEQ`

// BPL Branch on Result Plus
const BPL = `BPL`

// BMI Branch on Result Minus
const BMI = `BMI`

// BVC Branch on Overflow Clear
const BVC = `BVC`

// BVS Branch on Overflow Set
const BVS = `BVS`

// TAX Transfer Accumulator to Index X
const TAX = `TAX`

// TXA Transfer Index X to Accumulator
const TXA = `TXA`

// TAY Transfer Accumulator to Index Y
const TAY = `TAY`

// TYA Transfer Index Y to Accumulator
const TYA = `TYA`

// TSX Transfer Stack Pointer to Index X
const TSX = `TSX`

// TXS Transfer Index X to Stack Pointer
const TXS = `TXS`

// PHA Push Accumulator on Stack
const PHA = `PHA`

// PLA Pull Accumulator from Stack
const PLA = `PLA`

// PHP Push Processor Status on Stack
const PHP = `PHP`

// PLP Pull Processor Status from Stack
const PLP = `PLP`

// JMP Jump to New Location
const JMP = `JMP`

// JSR Jump to New Location Saving Return Address
const JSR = `JSR`

// RTS Return from Subroutine
const RTS = `RTS`

// RTI Return from Interrupt
const RTI = `RTI`

// CLC Clear Carry Flag
const CLC = `CLC`

// SEC Set Carry Flag
const SEC = `SEC`

// CLD Clear Decimal Mode
const CLD = `CLD`

// SED Set Decimal Mode
const SED = `SED`

// CLI Clear Interrupt Disable Status
const CLI = `CLI`

// SEI Set Interrupt Disable Status
const SEI = `SEI`

// CLV Clear Overflow Flag
const CLV = `CLV`

// BRK Break
const BRK = `BRK`

// NOP No Operation
const NOP = `NOP`
