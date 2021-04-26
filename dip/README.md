# 依存関係逆転の原則（Dependency inversion principle）

## ngpattern
Stable（安定）処理がFlexible（柔軟）を使用する  
StableがFlexibleの実装にもろ依存している（＝Flexibleな修正の影響をもろ受ける）  

## okpattern
stable → IService(interface) ←Flexible として双方をインターフェイスに依存するようにする  
インターフェイスの実装をFlexibleにさせる  
Stable がFlexible に依存することはなくなり、両方がIService に依存するようになる