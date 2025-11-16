import tkinter as tk
from tkinter import filedialog, messagebox
import pandas as pd
import os

class XLSXtoCSVConverter:
    def __init__(self, root):
        self.root = root
        self.root.title("Conversor XLSX para CSV")
        self.root.geometry("520x340")
        self.root.configure(bg="#f4f4f4")
        self.root.resizable(False, False)

        self.file_path = None
        self.csv_path = None

        # ==== Título ====
        title = tk.Label(
            root,
            text="Conversor XLSX → CSV",
            font=("Helvetica", 18, "bold"),
            bg="#f4f4f4",
            fg="#2c3e50"
        )
        title.pack(pady=(25, 15))

        # ==== Botão Selecionar ====
        select_btn = tk.Button(
            root,
            text="Selecionar arquivo XLSX",
            command=self.select_file,
            font=("Helvetica", 12, "bold"),
            width=28,
            height=1,
            bg="#3498db",
            fg="white",
            activebackground="#2980b9",
            relief="flat",
            cursor="hand2"
        )
        select_btn.pack(pady=(10, 5))

        # ==== Label nome do arquivo ====
        self.label_file = tk.Label(
            root,
            text="Nenhum arquivo selecionado",
            font=("Helvetica", 10),
            bg="#f4f4f4",
            fg="gray"
        )
        self.label_file.pack(pady=(0, 15))

        # ==== Botão Converter ====
        convert_btn = tk.Button(
            root,
            text="Converter para CSV",
            command=self.convert_file,
            font=("Helvetica", 12, "bold"),
            width=28,
            height=1,
            bg="#27ae60",
            fg="white",
            activebackground="#1e8449",
            relief="flat",
            cursor="hand2"
        )
        convert_btn.pack(pady=(0, 15))

        # ==== Label de status ====
        self.label_result = tk.Label(
            root,
            text="",
            font=("Helvetica", 10),
            bg="#f4f4f4",
            fg="black",
            wraplength=480,
            justify="center"
        )
        self.label_result.pack(pady=(10, 5), padx=20)

        # Rodapé
        footer = tk.Label(
            root,
            text="© 2025 XLSX to CSV Converter",
            font=("Helvetica", 8),
            bg="#f4f4f4",
            fg="#7f8c8d"
        )
        footer.pack(side="bottom", pady=10)

    def select_file(self):
        """Seleciona o arquivo XLSX"""
        file_path = filedialog.askopenfilename(
            title="Selecione um arquivo XLSX",
            initialdir=os.path.expanduser("~"),
            filetypes=[("Planilhas Excel", "*.xlsx *.xls")]
        )
        if file_path:
            self.file_path = file_path
            self.label_file.config(text=os.path.basename(file_path), fg="#2c3e50")
            self.label_result.config(text="")

    def convert_file(self):
        """Converte o arquivo XLSX para CSV no mesmo diretório"""
        if not self.file_path:
            messagebox.showwarning("Aviso", "Selecione um arquivo XLSX primeiro.")
            return

        try:
            df = pd.read_excel(self.file_path)
            self.csv_path = os.path.splitext(self.file_path)[0] + ".csv"
            df.to_csv(self.csv_path, index=False, encoding="utf-8-sig")

            self.label_result.config(
                text=(
                    "✅ Arquivo convertido com sucesso!\n\n"
                    "O arquivo CSV foi salvo no mesmo diretório do arquivo original:\n\n"
                    f"{os.path.dirname(self.csv_path)}"
                ),
                fg="#27ae60"
            )

        except Exception as e:
            messagebox.showerror("Erro", f"Erro ao converter o arquivo:\n{str(e)}")

# Execução principal
if __name__ == "__main__":
    root = tk.Tk()
    app = XLSXtoCSVConverter(root)
    root.mainloop()
