/**
 * このファイルはShadcnのインストール時に生成されたもので、clsxとtailwind-mergeライブラリを
 * 組み合わせて使用するためのユーティリティ関数を提供します。clsxは複数のクラス名を動的に
 * 結合するため、tailwind-mergeはTailwind CSSのユーティリティクラスを最適化するために使用されます。
 *
 * 使用方法:
 * - `cn` 関数は、動的にクラス名を結合し最適化する際に使用します。
 * - 複数のクラス名（文字列、オブジェクト、配列など）を引数として受け取り、
 *   最終的なクラス名の文字列を返します。
 *
 * 例:
 * const className = cn('text-center', {'text-lg': large}, ['p-4']);
 *
 * この関数はReact、VueなどのJavaScriptフレームワーク内での使用を想定していますが、
 * 任意のJavaScriptプロジェクトで利用可能です。
 */
import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
